use futures::{future, Future, Stream};
use orbital_api::builder::{server, BuildDeleteRequest, BuildLogResponse, BuildSummary};
use tower_grpc::{Request, Response};

#[derive(Clone, Debug)]
pub struct OrbitalApi;

impl orbital_api::builder::server::BuildService for OrbitalApi {
    type StartBuildFuture = future::FutureResult<Response<BuildSummary>, tower_grpc::Status>;
    type StopBuildFuture = future::FutureResult<Response<BuildSummary>, tower_grpc::Status>;
    type GetBuildLogsFuture = future::FutureResult<Response<BuildLogResponse>, tower_grpc::Status>;
    type DeleteBuildFuture = future::FutureResult<Response<BuildSummary>, tower_grpc::Status>;

    fn start_build(
        &mut self,
        request: Request<orbital_api::builder::BuildStartRequest>,
    ) -> Self::StartBuildFuture {
        println!("It works!");
        unimplemented!();
    }

    fn stop_build(
        &mut self,
        request: Request<orbital_api::builder::BuildStopRequest>,
    ) -> Self::StopBuildFuture {
        unimplemented!();
    }

    fn get_build_logs(
        &mut self,
        request: Request<orbital_api::builder::BuildLogRequest>,
    ) -> Self::GetBuildLogsFuture {
        unimplemented!();
    }

    fn delete_build(
        &mut self,
        request: Request<orbital_api::builder::BuildDeleteRequest>,
    ) -> Self::DeleteBuildFuture {
        unimplemented!();
    }
}
