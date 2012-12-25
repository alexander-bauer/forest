
Notes
=====

[07:38] <   DuoNoxSol> | Alright, so the idea is to loosely tie together a bunch of grove instances
[07:38] <   DuoNoxSol> | say you have a project
[07:38] <   DuoNoxSol> | like distru
[07:38] <   DuoNoxSol> | in which there are, say, four developers
[07:38] <   DuoNoxSol> | there may be more or less; I'm not really sure at the moment
[07:39] <   DuoNoxSol> | but say they're all running grove instances
[07:39] <   DuoNoxSol> | but we don't have a central repository to push to; we all pull from eachothers' branches
[07:39] <   DuoNoxSol> | and we don't have a central issue tracker or anything
[07:39] <   DuoNoxSol> | we only have grove
[07:40] <   DuoNoxSol> | but we want a central server (federated, meaning anyone can run their own forest, but any forest can host more than one 
          repository)
[07:40] <   DuoNoxSol> | which will house a bare repository we can push to, an issue tracker, a wiki
[07:40] <   DuoNoxSol> | (might not want the wiki; it's all unicorny at the moment)
[07:40] <   DuoNoxSol> | so we set up a server
[07:40] <   DuoNoxSol> | and run forest on it
[07:41] <   DuoNoxSol> | and feed it a list of all four developers' grove instances
[07:41] <   DuoNoxSol> | it will then open a bare repository with a name and blank issue tracker
[07:41] <   DuoNoxSol> | it may take ssh keys, too; that remains to be seen
[07:41] <   DuoNoxSol> | and all four developers are able to push to it
[07:41] <   DuoNoxSol> | but the main point is that it will periodically query their grove instances
[07:42] <   DuoNoxSol> | their particular repositories
[07:42] <   DuoNoxSol> | and in doing so collect information such as their recent commits
[07:42] <   DuoNoxSol> | their feature branches
[07:43] <   DuoNoxSol> | their repository's differences to the central server
[07:43] <   DuoNoxSol> | and it will then be able to display a page summarizing the activities of all of the developers
[07:43] <   DuoNoxSol> | as well as offer a central repository to clone from
[07:43] <   DuoNoxSol> | and issue tracker
