package webpb

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var ErrWebProgressBarNotInited = errors.New("WebProgressBar not initialized. It's important to use NewWebProgressBar()")

type WebProgressBar struct {
	init bool
	mu   sync.Mutex
	srv  http.Server

	pb      map[string]*ProgressBar
	pbOrder []string

	stopped bool
}

func NewWebProgressBar(addr string) *WebProgressBar {
	wpb := &WebProgressBar{
		pb:      map[string]*ProgressBar{},
		pbOrder: []string{},
	}

	handler := mux.NewRouter()
	handler.HandleFunc("/updates", func(w http.ResponseWriter, r *http.Request) {
		wpb.render(w, r)
	}).Methods(http.MethodGet)
	handler.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		wpb.refresh(w, r)
	}).Methods(http.MethodGet)
	handler.HandleFunc("/ui", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.New("ui").Parse(string(wpb.ui()))
		_ = tmpl.Execute(w, map[string]interface{}{"addr": addr})
	})

	wpb.srv = http.Server{Addr: addr, Handler: handler}
	wpb.init = true
	return wpb
}

func (b *WebProgressBar) initialized() bool {
	if b.pb == nil {
		return false
	}
	return true
}

func (b *WebProgressBar) AddProgressBar(name string, val, max int) (bar *ProgressBar, err error) {
	if !b.initialized() {
		err = ErrWebProgressBarNotInited
		return
	}

	if b.pb[name] != nil {
		err = fmt.Errorf("ProgressBar '%s' already exists", name)
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	tmpPtr := NewProgressBar(val, max)
	b.pb[name] = tmpPtr
	bar = tmpPtr
	b.pbOrder = append(b.pbOrder, name)
	return
}

func (b *WebProgressBar) GetProgressBar(name string) (bar *ProgressBar, err error) {
	if b.pb[name] == nil {
		err = fmt.Errorf("ProgressBar with name '%s' not exists", name)
		return
	}

	bar = b.pb[name]
	return
}

func (b *WebProgressBar) Run() {
	if !b.initialized() {
		return
	}

	go func() {
		if err := b.srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	fmt.Printf("ProgressBar UI: http://localhost%s/ui\n", b.srv.Addr)
}

func (b *WebProgressBar) Stop() {
	b.srv.Close()
}

func (b *WebProgressBar) refresh(w http.ResponseWriter, _ *http.Request) {
	data, _ := json.Marshal(b)
	w.Write(data)
}

func (b *WebProgressBar) render(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	var conn *websocket.Conn
	var err error
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	defer conn.Close()

	for {
		time.Sleep(time.Millisecond * 16) // Чуть больше 60Гц (~62,5)
		if b.updated() {
			if err = conn.WriteJSON(b); err != nil {
				return
			}
		}
	}
}

func (b *WebProgressBar) updated() (updated bool) {
	for k := range b.pb {
		if b.pb[k].Updated() {
			updated = true
		}
	}
	return
}

func (b *WebProgressBar) MarshalJSON() ([]byte, error) {
	tmp := map[string]interface{}{}

	var pbs []*ProgressBarView
	for i := range b.pbOrder {
		pb := b.pb[b.pbOrder[i]]

		pbs = append(pbs, &ProgressBarView{
			Name: b.pbOrder[i],
			Val:  pb.val,
			Max:  pb.max,
		})
	}

	tmp["progressBars"] = pbs
	return json.Marshal(tmp)
}

type ProgressBarView struct {
	Name  string `json:"Name"`
	Val   int    `json:"Val"`
	Max   int    `json:"Max"`
	Speed int    `json:"Speed"`
}

func (b *WebProgressBar) ui() []byte {
	return _uiTmpl
}
