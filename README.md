# Kontext

Get context information from files hidden your file tree.

Think of this as a different direnv that does not put anything into env and where every value is in a separate file.

## Installation

```[bash]
go install -ldflags="-s -w"
```

## Usage

```[bash]
# this will fail b/c there is no context file yet
kontext get project

# create a context file with a uuid in it
# (you cannot yet use the set-command for setting context one directory up)
uuidgen > ../.project.kontext

# this will show the path of the context file
kontext get project -f

# this will output the generated context uuid
kontext get project
```

### Set some context

```[bash]
kontext set project "$(uuidgen -7)"
```

### Get context file

```[bash]
kontext get project -f
```

### Get context directory

```[bash]
dirname "$(kontext get project -f)"
```

### Use a context with a JSON store

```[bash]
jq -n --arg date "$(date -u +"%Y-%m-%dT%H:%M:%SZ")" \
      --arg uuid "$(uuidgen)" \
      '{created: $date, uuid: $uuid}' > .project.kontext

jq -r '.uuid' "$(kontext get project -f)"
jq -r '.created' "$(kontext get project -f)"
```

### Use context as a taskwarrior project

You can use kontext to achieve context sensible taskwarrior tasks.

```[bash]
uuidgen > .project.kontext
task add project:$(kontext get project) "this is a new task for THIS project"
task add project:$(kontext get project) "this is another new task for THIS project"

# list all tasks for this project
task project:$(kontext get project)
```

Get wild and share your use case! <3
