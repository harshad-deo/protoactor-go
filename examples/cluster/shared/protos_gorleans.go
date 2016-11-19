// Code generated by protoc-gen-gogo.
// source: protos.proto
// DO NOT EDIT!

/*
Package shared is a generated protocol buffer package.

It is generated from these files:
	protos.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
	AddRequest
	AddResponse
*/
package shared

import log "log"
import github_com_AsynkronIT_gam_cluster_grains "github.com/AsynkronIT/gam/cluster/grains"
import github_com_AsynkronIT_gam_cluster "github.com/AsynkronIT/gam/cluster"
import github_com_AsynkronIT_gam_actor "github.com/AsynkronIT/gam/actor"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var xHelloFactory func() Hello

func HelloFactory(factory func() Hello) {
	xHelloFactory = factory
}

func GetHelloGrain(id string) *HelloGrain {
	return &HelloGrain{}
}

type Hello interface {
	SayHello(*HelloRequest) *HelloResponse
	Add(*AddRequest) *AddResponse
}
type HelloGrain struct {
	github_com_AsynkronIT_gam_cluster_grains.GrainMixin
}

func (g *HelloGrain) SayHello(r *HelloRequest) *HelloResponse {
	pid := github_com_AsynkronIT_gam_cluster.Get(g.Id(), "Hello")
	bytes, _ := proto.Marshal(r)
	gr := &github_com_AsynkronIT_gam_cluster_grains.GrainRequest{Method: "SayHello", MessageData: bytes}
	r0, _ := pid.AskFuture(gr, 1000)
	r1, _ := r0.Result()
	r2, _ := r1.(*github_com_AsynkronIT_gam_cluster_grains.GrainResponse)
	r3 := &HelloResponse{}
	proto.Unmarshal(r2.MessageData, r3)
	return r3
}

func (g *HelloGrain) SayHelloChan(r *HelloRequest) <-chan *HelloResponse {
	c := make(chan *HelloResponse, 1)
	defer close(c)
	c <- g.SayHello(r)
	return c
}

func (g *HelloGrain) Add(r *AddRequest) *AddResponse {
	pid := github_com_AsynkronIT_gam_cluster.Get(g.Id(), "Hello")
	bytes, _ := proto.Marshal(r)
	gr := &github_com_AsynkronIT_gam_cluster_grains.GrainRequest{Method: "Add", MessageData: bytes}
	r0, _ := pid.AskFuture(gr, 1000)
	r1, _ := r0.Result()
	r2, _ := r1.(*github_com_AsynkronIT_gam_cluster_grains.GrainResponse)
	r3 := &AddResponse{}
	proto.Unmarshal(r2.MessageData, r3)
	return r3
}

func (g *HelloGrain) AddChan(r *AddRequest) <-chan *AddResponse {
	c := make(chan *AddResponse, 1)
	defer close(c)
	c <- g.Add(r)
	return c
}

type HelloActor struct {
	inner Hello
}

func (a *HelloActor) Receive(ctx github_com_AsynkronIT_gam_actor.Context) {
	switch msg := ctx.Message().(type) {
	case *github_com_AsynkronIT_gam_cluster_grains.GrainRequest:
		switch msg.Method {
		case "SayHello":
			req := &HelloRequest{}
			proto.Unmarshal(msg.MessageData, req)
			a.inner.SayHello(req)
		case "Add":
			req := &AddRequest{}
			proto.Unmarshal(msg.MessageData, req)
			a.inner.Add(req)
		}
	default:
		log.Printf("Unknown message %v", msg)
	}
}

func init() {
	github_com_AsynkronIT_gam_cluster.Register("Hello", github_com_AsynkronIT_gam_actor.FromProducer(func() github_com_AsynkronIT_gam_actor.Actor { return &HelloActor{} }))
}
