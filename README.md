# GSP (GIT stacked PR) CLI

This is a highly-opinionated CLI tool to sync git stacked branches/PRs to your remote repository.

## Install

### Using homebrew

```
brew tap Open-SL/tap
brew install Open-SL/tap/gsp
```

## Commands

### Create a stack

A stack is collection of branches for a given repository. It will hold all the branch names and orders.
Use `-p` or `--primaryBranch` to mention the primary branch name of the repository. Default value is `master`.
```
gsp cs [stack-name] [-p primary-branch-name]
```

example,
```
gsp cs demo -d main
```

### Add a Branch to a stack

Adds a branch into a stack.
You can specify the priority of the branch using the -o or --order option. For example, 
a branch with an order of n+1 is a chained branch to a branch with an order of n.
```
gsp ab [stack-name] [branch-name] [-o order-of-the-branch]
```

example,
```
gsp ab demo feature/demo-branch-1
gsp ab demo feature/demo-branch-2
```

### Sync branches in to the remote

Syncs branches with remote.
```
gsp ss [stack-name] [optional-branch-name-to-initiate]
```

example,
```
# Syncs branches starting from main. Note that optional branch name to initiate is not mentioned.
# This will
# pull main branch into the feature/demo-branch-1
# push feature/demo-branch-1 into the remote
# checkout and update feature/demo-branch-2 from feature/demo-branch-1
# push updated changes to feature/demo-branch-2's remote
gsp ss demo
```

Furthermore, gsp provided initializing a sync from desired branch.
if a stack contains 3 branches and only the 2nd branch was updated,
following commands will start the sync starting from the second branch.

```
# Adds the third branch
gsp ab demo feature/demo-branch-3

# Starts sync from the second branch
gsp ss demo feature/demo-branch-2
```

### List stacks

Lists stack names.
```
gsp ls
```

### Delete stack

Deletes a named stack from the gsp stack list.

example,
```
gsp ds demo
```

in addition to that, users can directly update/check the data store at `~/.gsp/data.json` at anytime.

### Development

gsp uses [asdf](https://asdf-vm.com/) to manage `golang` version.

Use
```
asdf install
```
to install the required golang version.

Use
```
go build & go install
```
to tryout the CLI in local.

## Future Developments

- List branches in a selected stack - gsp lb stack-name
- Remove a branch from a selected stack - gsp rb -s stack-name branch-name
