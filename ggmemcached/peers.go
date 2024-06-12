package ggmemcached

import pb "ggmemcached/ggmemcachedpb"

// PeerPicker 根据传入的 key 选择相应节点
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter 从应 group 查找对缓存值
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error //
}
