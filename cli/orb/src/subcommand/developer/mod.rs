use structopt::StructOpt;

use crate::{GlobalOption, SubcommandError};

pub mod git;
pub mod docker;

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub enum DeveloperType {
    Git(git::SubcommandOption),
    Docker(docker::SubcommandOption),
}

pub fn subcommand_handler(
    global_option: GlobalOption,
    dev_subcommand: DeveloperType,
) -> Result<(), SubcommandError> {
    match dev_subcommand {
        DeveloperType::Git(sub_option) => git::subcommand_handler(global_option, sub_option),
        DeveloperType::Docker(sub_option) => docker::subcommand_handler(global_option, sub_option),
    }
}