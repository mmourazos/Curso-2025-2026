# Checklist de instalación y configuración de Neovim

## Instalación de Neovim

Si tenemos Windows y somos administradores: `winget install neovim`.

Si no somos administradores usaremos [scoop](https://scoop.sh/).

### Instalación de `scoop`

Una vez instalado scoop:

```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
```

Instalamos `neovim` con `scoop install neovim`.

## Instalación de `lazyvim`

Para instalar [lazyvim](https://www.lazyvim.org/) hemos de consultar las dependencias:

* Neovim >= 0.9.0 - :smile:
* Git >= 2.19.0 (for partial clones support) - ✔️
* a Nerd Font(v3.0 or greater)

Par instalar _Nerd Fonts_ hemos de añadir el _bucket_ a scoop: `scoop bucket add nerd-fonts`. Para instalar una fuente pondremos el nombre de la fuente `scoop search fantasque`, por ejemplo, y luego seleccionamos el nombre correcto `scoop install FantasqueSansMono-NF`.

* a C compiler for nvim-treesitter : `scoop install gcc`.
* curl for blink.cmp (completion engine) : `scoop install curl`.
* for fzf-lua (optional)
  * fzf: fzf (v0.25.1 or greater)
  * live grep: ripgrep
  * find files: fd
* a terminal that support true color and undercurl:
  * wezterm (Linux, Macos & Windows)

Wezterm, a día de hoy (9-9-2025) da problemas en Windows para escribir _underculr_. Para evitar que no se nos muestren los errores ortográficos y nos volvamos locos instalaremos el _bucket_ de versiones de scoop: `scoop bucket add versions` y a continuación podremos instalar la versión _nightly_ de Wezterm (que sí soporta los _undercurl_) `scoop install wezterm-nightly`.

En resumen, para tener todo lo básico de lazyvim:

```powershell
scoop add bucket nerd-fonts
scoop install FantasqueSansMono-NF
scoop install gcc
scoop install curl
scoop install fzf
scoop install ripgrep
scoop install fd
scoop bucker add versions
scoop install wezterm-nightly
```

Hemos de configurar el terminal para que use nuestra fuente:

En el archivo de configuración de wezterm: `~/.wezterm.lua`

```lua
config.font = wezterm.font("FantasqueSansM Nerd Font")
```

Llegados a este punto podríamos dar por instalado lazyvim.

## Detector de errores ortográficos

Para que neovim detecte errores ortográficos:

En el archivo de configuración de neovim (dentro de `$env:LOCALAPPDATA`): `./nvim/lua/config/options.lua`, añadiremos:

```lua
-- Spelling
vim.o.spelllang = "es,gl,en_gb,en_us"
vim.o.spell = true
```

Si reiniciamos neovim, éste detectará que faltan los archivos necesarios para la corrección ortográfica e intentará descargarlos. Para que funcione hemos de tener configurado neovim par que use `cmd.exe` como shell (está así por defecto) e instalar `weget`: `scoop install wget`.

Una vez instalado `weget` y reiniciado neovim se instalarán los ficheros necesarios y tendremos configurada la corrección ortográfica.

## Actualización de la shell

Para instalar la versión más reciente de powershell: `scoop install pwsh`.

## Para ponerla como shell por defecto de neovim

Añadir a `nvim/lua/config/options.lua`:

```lua
-- Usar powershell como shell del terminal
vim.o.shell = "pwsh"
```

Si da problemas podría ser necesario añadir también:

```lua
vim.o.shellcmdflag =
  "-NoLogo -NoProfile -ExecutionPolicy RemoteSigned -Command [Console]::InputEncoding=[Console]::OutputEncoding=[System.Text.Encoding]::UTF8;"
vim.o.shellredir = "2>&1 | Out-File -Encoding UTF8 %s; exit $LastExitCode"
vim.o.shellpipe = "2>&1 | Out-File -Encoding UTF8 %s; exit $LastExitCode"
vim.o.shellquote = ""
vim.o.shellxquote = ""
```

