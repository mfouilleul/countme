package main

import (
	"fmt"
	"net/http"

	"github.com/cenkalti/backoff"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type app struct {
	Port  int
	Redis db `json:"database"`
}

type db struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Conn     *redis.Client
}

func (a *app) Help(w http.ResponseWriter, r *http.Request) {
	help := `PUT /incr
GET /count
GET /version`

	fmt.Fprintf(w, help)
}

func (a *app) Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, version)
}

func (a *app) Count(w http.ResponseWriter, r *http.Request) {

	val, err := a.Redis.Conn.Get("countme").Result()
	if err != nil {
		fmt.Fprintf(w, "error")
	}

	fmt.Fprintf(w, val)
}

func (a *app) Incr(w http.ResponseWriter, r *http.Request) {
	_, err := a.Redis.Conn.Incr("countme").Result()

	if err != nil {
		fmt.Fprintf(w, "error")
	}

	a.Count(w, r)
}

func (a *app) Initialize() error {
	cs := fmt.Sprintf("%s:%d", a.Redis.Hostname, a.Redis.Port)
	a.Redis.Conn = redis.NewClient(&redis.Options{
		Addr: cs,
	})

	pingRedis := func() error {
		_, err := a.Redis.Conn.Ping().Result()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"hostname": a.Redis.Hostname,
				"port":     a.Redis.Port,
			}).Info("Redis Connection failed")
			return err
		}
		return nil
	}

	b := backoff.NewExponentialBackOff()
	err := backoff.Retry(pingRedis, b)

	return err
}

func (a *app) Run() error {
	http.HandleFunc("/", a.Help)
	http.HandleFunc("/count", a.Count)
	http.HandleFunc("/incr", a.Incr)
	http.HandleFunc("/version", a.Version)

	addr := fmt.Sprintf(":%d", a.Port)
	return http.ListenAndServe(addr, nil)
}
