# cred-assist

This application aims to provide a simple means of managing a small set of AWS credentials (access key, secret key, and MFA) with a larger set of assumable roles.  The intent is to make it easy to keep base credentials minimized (as in: don't have a set for every little thing) and encourage secure usage (as in: use MFA, plus don't have tons of keys floating around as in the previous item).  Doing this really requires assuming roles, which can be a bit of a hassle in a local dev environment but has the benefit of more realistically mimicking how this will run in AWS if you're using roles and profiles (and you **are** doing that, *right?*).

## about

This is my first Go application, and will likely also be my first for [whatever UI framework I end up using].  I'm hoping to learn something while also creating a tool that would be useful to myself and my co-workers in our daily work.  I'm hoping to do things as right as I can given my current understanding, but I have no doubt I'll learn things along the way that change that understanding and make me realize I did things totally wrong.  As such, I'm totally open to feedback, assistance, etc.