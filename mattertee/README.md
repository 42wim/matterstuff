# mattertee #

*mattertee* works like [tee](http://en.wikipedia.org/wiki/Tee_(command)) command (and inspired by [slacktee](https://github.com/course-hero/slacktee)
Instead of writing the standard input to files, *mattertee* posts it to your mattermost installation


## binaries
Binaries will be found [here] (https://github.com/42wim/matterstuff/releases/)

## building
Make sure you have [Go](https://golang.org/doc/install) properly installed, including setting up your [GOPATH] (https://golang.org/doc/code.html#GOPATH)

```
cd $GOPATH
go get github.com/42wim/matterstuff/mattertee
```

You should now have mattertee binary in the bin directory:

```
$ ls bin/
mattertee
```

## usage

### On the command line

```
Usage of ./mattertee:
  -c string
         Post input values to specified channel or user.
  -i string
        This url is used as icon for posting.
  -l string
        Specify the language used for syntax highlighting (ruby/python/...).
  -m string
        Mattermost incoming webhooks URL.
  -n    Post input values without buffering.
  -p    Don't surround the post with triple backticks.
  -u string
        This username is used for posting. (default "mattertee")
  -x    Add extra info (user/hostname/timestamp).
```

### Configuration files

Mattertee will also read from configuration files in order. Later files can (partially) override earlier files.

On Linux: /etc/mattertee.conf, $HOME/.mattertee.conf

On Windows: $HOMEDIR/.mattertee.conf

On all OSes, a file '.mattertee.conf' in the current directory will be read too.

The files (if present) should be yaml-formatted; eg.:

```yaml
username: thisisme
channel: mychannel
icon_url: http://..../image.png
matter_url: http://mattermost.at.my.domain/hooks/hookid
title: Some title
language: ruby
plain_text: true
no_buffer: false
extra: true
```

Command line options will still override anything set in the configuration files.

### Using environment variables

You can also set MM_HOOK containing your mattermost incoming webhook URL as an enviroment variable.

```
export MM_HOOK=https://yourmattermost/hooks/webhookkey
```

### example
```
uptime | mattertee -c off-topic -x -m https://yourmattermost/hooks/webhookkey
```

or if you've set MM_HOOK environment variable:

```
uptime | mattertee -c off-topic -x
```

![Image](http://snag.gy/qJomi.jpg)

```
cat test.rb |mattertee -c off-topic -l ruby
```

![Image](http://snag.gy/58ryr.jpg)

## examples (taken from slacktee)
If you'd like to post the output of `ls` command, you can do it like this.

```
ls | mattertee
```

To post the output of `tail -f` command line by line, use `-n` option.

```
tail -f foobar.log | mattertee -n
```

You can specify `channel`, `username`, `iconurl` too.

```
ls | mattertee -c "general" -u "mattertee" -i "http://myimage.png"
```

Of course, you can connect another command with pipe.

```
ls | mattertee | email "ls" foo@example.com
```
