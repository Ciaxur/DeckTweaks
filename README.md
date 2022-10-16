# DeckTweaks
A [decky-loader](https://github.com/SteamDeckHomebrew/deckly-loader) homebrew plugin that contains configurable SteamDeck tweaks.

## Features
### Battery Monitor
Backend API can monitor the battery percentage, providing feedback to the user (*through toast notification*) upon low battery charge
or a given max/min charge percentages.

## Backend API
The backend server is written in go and can be built by invoking the [build.sh](backend/build.sh) script and then starting the server
manually using:
```sh
> server start
2022/10/15 17:53:18 Listening on localhost:3001.
```

<p align="center">
   <img src="docs/backend_api.svg", width=80% >
   </img>
</p>

# License
Licensed under [GNU GPLv3](LICENSE).