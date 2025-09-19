# Soluciones

## treesitter no compila

Solución al error de que treesitter no compile nada:

> I had a similar issue "error: linking with link.exe failed: exit code: 1"
>
> To solve it, I did
>
> rustup toolchain install stable-x86_64-pc-windows-gnu
> then
>
> rustup default stable-x86_64-pc-windows-gnu

## Copilot no hace nada

Es posible que esté relacionado con Go???

Creo que cuando "reinstalo" nvim/lazyvim funciona. Cuando activo (o se activan automáticamente) cosas de LazyExtras o de Mason algo _se rompe_ y, aunque dice que todo correcto, no puedo ver las sugerencias.

Veamos lo que acabo de hacer:

1. Instalación desde cero de LazyVim.
2. Modifiqué el archivo de configuración `lua/config/options.lua` para añadir mi configuración de _spell_.
   - No se descargaron los ficheros de idiomas pero la ortografía en español e inglés está funcionando ¿¿¿???
3. Activé **Copilot** en LazyExtras.

En este punto, copilot funciona perfectamente.

Haré los siguientes pasos uno a uno para ver en qué punto deja de funcionar:

1. Activé en LazyExtras el lenguaje `Markdown` y reinicié `nvim` (repetiré el reinicio cada vez que haga un cambio).
   - Copilot sigue funcionando.
2. Añado `nvim-lint.lua` para eliminar el warning de `line-length`.
   - Copilot sigue funcionando.
3. Como creo que no debería influir para nada instalo también `Copilot-chat`.
   - Copilot sigue funcionando.
4. Voy a instalar soporte para `Go`.
   - Copilot sigue funcionando.
5. Instalo `emmet-language-server`.
   - Copilot sigue funcionando.
6. Vuelvo a instalar cosas de mi configuración personal.
   1. Autosave.
   2. Instalo temas.
   3. Instalo `Themery`.

- **Copilot ha dejado de funcionar**.

Creo que es algo temporal por la forma en que van apareciendo más sugirencias de las `abc` y `[T]`.

Todo lo anterior lo estaba ejecutando desde `cmd.exe` y **es posible** que copilot haya dejado de funcionar cuando empecé a lanzar nvim desde otra shell.

Me salen sugerencia `abc` que me tocan los huevos pero nada de copilot.
