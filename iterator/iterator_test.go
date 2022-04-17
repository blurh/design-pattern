package iterator

import (
    "reflect"
    "testing"
)

func TestDesign(t *testing.T) {
    t.Run("test for iterator", func(t *testing.T) {
        tree := InitTree(10)
        tree.LeftAdd(5)
        tree.RightAdd(13)
        if tree.Index != 10 || tree.LeftNode.Index != 5 || tree.RightNode.Index != 13 {
            t.Error("tree test fail")
        }

        tree.LeftNode.RightAdd(4)
        tree.LeftNode.LeftAdd(3)
        tree.RightNode.LeftAdd(11)
        //          10
        //         /  \
        //        5    13
        //       / \   /
        //      3   4 11

        gotBreadthList := tree.BreadthFirstSearch()
        wantBreathList := []int{10, 5, 13, 3, 4, 11}
        if !reflect.DeepEqual(wantBreathList, gotBreadthList) {
            t.Error("test breadth fail")
        }
        gotDepthList := tree.DepthFirstSearch()
        wantDepthLIst := []int{10, 5, 3, 4, 13, 11}
        if !reflect.DeepEqual(wantDepthLIst, gotDepthList) {
            t.Error("test depth fail")
        }
    })
}
