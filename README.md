<h1 align="center">
  GoNamer
  <br>
  <i>A Media Renaming Tool</i>
  <p>
    <img src="https://imgs.search.brave.com/811XP1RwmaV1BGXmXM8o3m8PYtEKXs9EzvN3qxl8YHE/rs:fit:860:0:0/g:ce/aHR0cHM6Ly9nb3Bo/ZXJzb3VyY2UuY29t/L2ltZy9taWMtZHJv/cC5wbmc" width="200" height="200" />
  </p>
</h1>

<p align="center">
  <img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/palette/macchiato.png" width="400" />
</p>

## What is it

I wanted to practice writing Go, and used to use <a href="https://github.com/jkwill87/mnamer">Mnamer</a>
for renaming all of my media. 

This project is based on that. There are still some things
that need to be done such as:

- [ ] Integrating with theTVDB for episodes
- [ ] Adding Charm options for selection
- [x] Reduce boilerplate by having it rename once even if more than one file
- [ ] Throw an exception if the regex for movies is not found (no year)
- [ ] Parse TV shows based on season and episode
- [ ] Pull TV episode names

So far it works well on movies
