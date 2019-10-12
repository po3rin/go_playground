package main

import (
	"fmt"
	"log"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func main() {
	s := op.NewScope()
	c := op.Const(s, "Hello form TensorFlow version"+tf.Version())
	graph, err := s.Finalize()
	if err != nil {
		log.Fatal(err)
	}

	sess, err := tf.NewSession(graph, nil)
	if err != nil {
		log.Fatal(err)
	}
	output, err := sess.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output[0].Value())
}
