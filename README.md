# YAC = Yet Another Cron

This is a super-simpliffied cron, that run tasks every `n` minutes (where `n`) is defined in a configuration JSON file...

## Install

- From sources
    ```bash
    go get -u -v github.com/skryde/yac
    $GOPATH/bin/yac  # run yac.
    ```

- From compiled release
    ```bash
    # Coming soon
    ```

## Files

- `crons.json`:

    In Windows it's saved in `%APPDATA%\yac`

    In *nix it's saved in `$HOME/.config/yac`

- `yac.log`

    In Windows it's saved in `%APPDATA%\yac\logs`
    
    In *nix it's saved in `/var/log`

## FAQ

**Q:** Is this simple 'cron' truly necessary?

**A:** No, it isn't at all; but i was bored, wanted to write some code, keep learning Go and needed a 'cron' implementation for Windows.
