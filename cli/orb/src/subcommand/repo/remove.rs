use structopt::StructOpt;

use crate::{repo::SubcommandOption, GlobalOption, SubcommandError};

use orbital_headers::code::{client::CodeServiceClient, GitRepoRemoveRequest};
use orbital_services::ORB_DEFAULT_URI;
use tonic::Request;

use log::debug;
use std::path::PathBuf;

#[derive(Debug, StructOpt, Clone)]
#[structopt(rename_all = "kebab_case")]
pub struct ActionOption {
    /// Repo path
    #[structopt(parse(from_os_str), env = "PWD")]
    path: PathBuf,

    /// Name of Orbital org
    #[structopt(long, env = "ORB_DEFAULT_ORG")]
    org: Option<String>,
}

pub async fn action_handler(
    _global_option: GlobalOption,
    _subcommand_option: SubcommandOption,
    action_option: ActionOption,
) -> Result<(), SubcommandError> {
    let mut client = CodeServiceClient::connect(format!("http://{}", ORB_DEFAULT_URI)).await?;

    let request = Request::new(GitRepoRemoveRequest {
        org: action_option.org.unwrap_or_default(),
        ..Default::default()
    });

    debug!("Request for git repo remove: {:?}", &request);

    let response = client.git_repo_remove(request).await?;
    println!("RESPONSE = {:?}", response);
    Ok(())
}
