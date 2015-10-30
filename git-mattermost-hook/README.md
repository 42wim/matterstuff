# Git post-receive hook for mattermost 
(Based on https://github.com/chriseldredge/git-slack-hook)

This is a bash script that posts a message into your [Mattermost](https://mattermost.org) channel when changes are pushed.

Hook this script into `post-receive` for your git repositories.

## How to Install
Requires [mattertee](https://github.com/42wim/matterstuff/tree/master/mattertee)

Note: some git repositories may be "bare". You'll know if your repo is bare or not by checking for a `.git` folder where your repo lives.

https://raw.githubusercontent.com/42wim/matterstuff/master/mattertee/README.md
Download [git-mattermost-hook](https://raw.githubusercontent.com/42wim/matterstuff/master/git-mattermost-hook/git-mattermost-hook) onto the server which hosts your git repo.

For bare repos, copy/rename it as `/path/to/your/repo/hooks/post-receive`.

For normal/non-bare repos, copy/rename it as `/path/to/your/repo/.git/hooks/post-receive`.

Finally, `chmod +x post-receive` to allow the script to be executed.

## Configuration

Add an Incoming WebHooks integration in your mattermost
Make note of the webhook URL.

    git config hooks.slack.webhook-url 'https://yourmattermost/hooks/key'

Specify the path of [mattertee](https://github.com/42wim/matterstuff/tree/master/mattertee) binary
    git config hooks.slack.mattertee '/bin/mattertee'

## Optional

    git config hooks.slack.channel '#general'

        '#channelname' - post to channel
        'groupname' - post to group

Specifies a channel to post in Slack instead of the default.

    git config hooks.slack.username 'git'

Specifies a username to post as. If not specified, the default name `incoming-webhook` will be used.

    git config hooks.slack.icon-url 'https://example.com/icon.png'

Specifies an emoji icon to display in Slack instead of the default.

    git config hooks.slack.repo-nice-name 'My Awesome Repository'

Specifies a repository nice name that will be shown in messages.

    git config hooks.slack.show-only-last-commit true

Specifies whether you want to show only the last commit (or all) when pushing multiple commits.

    git config hooks.slack.branch-regexp regexp

Specifies if you want to send only certain branches

## Linking to Changesets

When the following parameters are set, revision hashes will be turned into links to a web view of your repository.

    git config hooks.slack.repos-root '/path/to/repos'
    git config hooks.slack.changeset-url-pattern 'http://yourserver/%repo_path%/changeset/%rev_hash%'

For example, if your repository is in `/usr/local/repos/myrepo`, set repos_root to `/usr/local/repos/` and set `changeset_url_pattern` to `http://yourserver/%repo_path%/changeset/%rev_hash%` or whatever.

Links can also be created that summarize a list of commits:

    git config hooks.slack.compare-url-pattern 'http://yourserver/%repo_path%/compare/%old_rev_hash%..%new_rev_hash%'
