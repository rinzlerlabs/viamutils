package api

import (
	"context"

	agent_proto "go.viam.com/api/app/agent/v1"
	build_proto "go.viam.com/api/app/build/v1"
	cloudslam_proto "go.viam.com/api/app/cloudslam/v1"
	data_proto "go.viam.com/api/app/data/v1"
	dataset_proto "go.viam.com/api/app/dataset/v1"
	datasync_proto "go.viam.com/api/app/datasync/v1"
	mlinference_proto "go.viam.com/api/app/mlinference/v1"
	mltraining_proto "go.viam.com/api/app/mltraining/v1"
	model_proto "go.viam.com/api/app/model/v1"
	app_proto "go.viam.com/api/app/v1"
	"go.viam.com/rdk/logging"
	"go.viam.com/utils/rpc"
)

func NewAppClientFromApiCredentials(ctx context.Context, logger logging.Logger, apiKeyName string, apiKey string) (app_proto.AppServiceClient, error) {
	conn, err := rpc.DialDirectGRPC(
		ctx,
		"app.viam.com:443",
		logger,
		rpc.WithEntityCredentials(
			apiKeyName,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiKey,
			}),
	)
	if err != nil {
		return nil, err
	}

	return app_proto.NewAppServiceClient(conn), nil
}

func NewAgentClientFromApiCredentials(ctx context.Context, logger logging.Logger, apiKeyName string, apiKey string) (agent_proto.AgentDeviceServiceClient, error) {
	conn, err := rpc.DialDirectGRPC(
		ctx,
		"app.viam.com:443",
		logger,
		rpc.WithEntityCredentials(
			apiKeyName,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiKey,
			}),
	)
	if err != nil {
		return nil, err
	}

	return agent_proto.NewAgentDeviceServiceClient(conn), nil
}

func NewBuildClientFromApiCredentials(ctx context.Context, logger logging.Logger, apiKeyName string, apiKey string) (build_proto.BuildServiceClient, error) {
	conn, err := rpc.DialDirectGRPC(
		ctx,
		"app.viam.com:443",
		logger,
		rpc.WithEntityCredentials(
			apiKeyName,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiKey,
			}),
	)
	if err != nil {
		return nil, err
	}

	return build_proto.NewBuildServiceClient(conn), nil
}

func NewCloudSLAMClientFromApiCredentials(ctx context.Context, logger logging.Logger, apiKeyName string, apiKey string) (cloudslam_proto.CloudSLAMServiceClient, error) {
	conn, err := rpc.DialDirectGRPC(
		ctx,
		"app.viam.com:443",
		logger,
		rpc.WithEntityCredentials(
			apiKeyName,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiKey,
			}),
	)
	if err != nil {
		return nil, err
	}

	return cloudslam_proto.NewCloudSLAMServiceClient(conn), nil
}

func NewDataClientFromApiCredentials(ctx context.Context, logger logging.Logger, apiKeyName string, apiKey string) (data_proto.DataServiceClient, error) {
	conn, err := rpc.DialDirectGRPC(
		ctx,
		"data.viam.com:443",
		logger,
		rpc.WithEntityCredentials(
			apiKeyName,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiKey,
			}),
	)
	if err != nil {
		return nil, err
	}

	return data_proto.NewDataServiceClient(conn), nil
}

func NewDataSetClientFromApiCredentials(ctx context.Context, logger logging.Logger, apiKeyName string, apiKey string) (dataset_proto.DatasetServiceClient, error) {
	conn, err := rpc.DialDirectGRPC(
		ctx,
		"data.viam.com:443",
		logger,
		rpc.WithEntityCredentials(
			apiKeyName,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiKey,
			}),
	)
	if err != nil {
		return nil, err
	}

	return dataset_proto.NewDatasetServiceClient(conn), nil
}

func NewDataSyncClientFromApiCredentials(ctx context.Context, logger logging.Logger, apiKeyName string, apiKey string) (datasync_proto.DataSyncServiceClient, error) {
	conn, err := rpc.DialDirectGRPC(
		ctx,
		"data.viam.com:443",
		logger,
		rpc.WithEntityCredentials(
			apiKeyName,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiKey,
			}),
	)
	if err != nil {
		return nil, err
	}

	return datasync_proto.NewDataSyncServiceClient(conn), nil
}

func NewMLInferenceClientFromApiCredentials(ctx context.Context, logger logging.Logger, apiKeyName string, apiKey string) (mlinference_proto.MLInferenceServiceClient, error) {
	conn, err := rpc.DialDirectGRPC(
		ctx,
		"app.viam.com:443",
		logger,
		rpc.WithEntityCredentials(
			apiKeyName,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiKey,
			}),
	)
	if err != nil {
		return nil, err
	}

	return mlinference_proto.NewMLInferenceServiceClient(conn), nil
}

func NewMLTrainingClientFromApiCredentials(ctx context.Context, logger logging.Logger, apiKeyName string, apiKey string) (mltraining_proto.MLTrainingServiceClient, error) {
	conn, err := rpc.DialDirectGRPC(
		ctx,
		"app.viam.com:443",
		logger,
		rpc.WithEntityCredentials(
			apiKeyName,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiKey,
			}),
	)
	if err != nil {
		return nil, err
	}

	return mltraining_proto.NewMLTrainingServiceClient(conn), nil
}

func NewModelClientFromApiCredentials(ctx context.Context, logger logging.Logger, apiKeyName string, apiKey string) (model_proto.ModelServiceClient, error) {
	conn, err := rpc.DialDirectGRPC(
		ctx,
		"app.viam.com:443",
		logger,
		rpc.WithEntityCredentials(
			apiKeyName,
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: apiKey,
			}),
	)
	if err != nil {
		return nil, err
	}

	return model_proto.NewModelServiceClient(conn), nil
}
