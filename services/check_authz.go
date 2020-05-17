package services

import (
	// "context"
	// "error"
	"io"

	pb "github.com/gitkado/go-grpc-study/pb"
)

type AuthorizationServerService struct {}

// gRPC Unary RPC
// func (s *AuthorizationServerService) CheckAuthz(ctx context.Context, msg *pb.CheckAuthzMessage) (*pb.CheckAuthzResponse, error) {
// 	switch msg.UserId {
// 	case "gitkado":
// 		return &pb.CheckAuthzResponse{
//             IsCheck: true,
//         }, nil
// 	default:
// 		return &pb.CheckAuthzResponse{
//             IsCheck: false,
// 		}, nil
// 	}
// 	// 書いてるだけ(switchでreturnされる)
// 	return nil, status.Errorf(codes.Unimplemented, "method CheckAuthz not implemented")
// }

// gRPC Bidirectional streaming RPC
func (s *AuthorizationServerService) CheckAuthz(stream pb.Authorization_CheckAuthzServer) error {
    for {
        msg, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
		}

		switch msg.UserId {
		case "gitkado":
			stream.Send(&pb.CheckAuthzResponse{
				IsCheck: true,
			})
		default:
			stream.Send(&pb.CheckAuthzResponse{
				IsCheck: false,
			})
		}
	}
	// RPC終了
    return nil
}
