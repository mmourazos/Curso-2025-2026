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
