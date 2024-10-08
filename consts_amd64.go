//go:build linux && amd64
// +build linux,amd64

package v4l2

// #include "linux/videodev2.h"
// #include "stdio.h"
//
// int main() {
//   printf("VIDIOC_DQBUF = 0x%x\n", VIDIOC_DQBUF);
//   printf("VIDIOC_QBUF = 0x%x\n", VIDIOC_QBUF);
//   printf("VIDIOC_QUERYBUF = 0x%x\n", VIDIOC_QUERYBUF);
//   printf("VIDIOC_G_EXT_CTRLS = 0x%x\n", VIDIOC_G_EXT_CTRLS);
//   printf("VIDIOC_S_EXT_CTRLS = 0x%x\n", VIDIOC_S_EXT_CTRLS);
//   printf("VIDIOC_S_FMT = 0x%x\n", VIDIOC_S_FMT);
//   return 0;
// }

const (
	VIDIOC_DQBUF       = 0xc0585611
	VIDIOC_QBUF        = 0xc058560f
	VIDIOC_QUERYBUF    = 0xc0585609
	VIDIOC_G_EXT_CTRLS = 0xc0205647
	VIDIOC_S_EXT_CTRLS = 0xc0205648
	VIDIOC_S_FMT       = 0xc0d05605
)
