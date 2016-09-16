# truncater
A very fast truncater of files identified by a pattern, including recursive subdirectories

Usage: ./truncater -d /home/example/ -p *.log

# Docker
By default, the truncate directory is /var/lib/docker/containers and the pattern is *-json.log. This cleans the default docker logs in a debian/ubuntu default installation.

So in case you want to use it as a docker log truncater:

Usage: ./truncater

will do the job.
