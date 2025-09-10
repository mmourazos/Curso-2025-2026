# Apuntes de markdown

## Origen de Markdown

El lenguaje de marcado Markdown fue creado por John Gruber en 2004, con la colaboración de Aaron Swartz. La idea principal detrás de Markdown era crear un lenguaje de marcado que fuera fácil de leer y escribir en su forma original, y que pudiera convertirse fácilmente en HTML para su publicación en la web.

Cuando se refieren a que es fácil de leer y escribir, significa que es fácil interpretar lo que se quiere expresar sin necesidad de ver la representación _renderizada_ (por ejemplo, en HTML). Esto es especialmente útil para escritores y desarrolladores que desean centrarse en el contenido sin preocuparse demasiado por la presentación.

Un ejemplo de código markdown (además de este propio texto) es el siguiente:

```markdown
# Título de nivel 1

Un párrafo de texto con **negrita** (`**negrita**`) y *cursiva* (`*cursiva*`) (también se puede usar `_` y  `__` en lugar de `*` pero **hay que ser consistente**).

## Sección/título de nivel 2

Una lista no ordenada:

* Hay que preceder.
* Cada elemento.
* De un `*`.

Una lista ordenada:

1. primer elemento.
2. segundo elemento.
3. tercer elemento.

No es necesario que los números sean consecutivos, el renderizador se encarga de eso.

etc.

```

## Estructura básica

La estructura básica de un documento en Markdown incluye:

* Los títulos y subtítulos, que se crean utilizando el símbolo `#` seguido de un espacio y el texto del título. El número de `#` indica el nivel del título (por ejemplo, `#` para el nivel 1, `##` para el nivel 2, etc.).
* Las listas, que pueden ser ordenadas (numeradas) o no ordenadas (con viñetas). Las listas no ordenadas se crean utilizando `*`, `-` o `+` seguido de un espacio y el texto del elemento de la lista. Las listas ordenadas se crean utilizando números seguidos de un punto y un espacio.
* Los párrafos, que se crean simplemente escribiendo el texto sin ningún símbolo especial. Los párrafos se separan por una línea en blanco.
* Las negritas y cursivas, que se crean utilizando `**` o `__` para negritas y `*` o `_` para cursivas.
* Los enlaces, que se crean utilizando el formato `[texto del enlace](URL)`.
* Las imágenes, que se crean utilizando el formato `![texto alternativo](URL de la imagen)`.

Además de estos elementos básicos, Markdown también soporta otros elementos como citas, bloques de código, tablas, etc.

## Conversión a HTML

### `Pandoc`

Pandoc es una herramienta de conversión de documentos que puede convertir archivos en formato Markdown a HTML y muchos otros formatos. Es una herramienta de línea de comandos que se puede instalar en la mayoría de los sistemas operativos.

Para instalar Pandoc en nuestro sistema:

* En Windows (una vez instalado [scoop](https://scoop.sh/)): `scoop install pandoc`.
* En Linux (Debian/Ubuntu): `sudo apt install pandoc`.

Para convertir un archivo Markdown a HTML, se puede utilizar el siguiente comando en la terminal:

```bash
pandoc --css=styling.css -s -f markdown+smart --toc --metadata --pagetitle="awesome ldap" --to=html5 README.md -o index.html
```

Analicemos el comando anterior:

* `pandoc`: es el comando para ejecutar la herramienta Pandoc.
* `--css=styling.css`: especifica una hoja de estilos CSS para aplicar al fichero HTML generado.
* `-s`: indica que se debe generar un documento HTML completo (con encabezado y cuerpo).
* `-f markdown+smart`: especifica el formato de entrada, en este caso Markdown.
* `--pagetitle`: especifica el título de la página HTML.
* `--toc`: genera una tabla de contenidos basada en los títulos del documento.
* `--metadata`: permite añadir metadatos al documento HTML.
* `--to=html5`: especifica el formato de salida, en este caso HTML5.
* `-o index.html`: especifica el nombre del archivo de salida.
* `README.md`: es el archivo de entrada en formato Markdown.

Busquemos ahora una hoja de estilos CSS sencilla para aplicar al fichero HTML generado. Uno de los primeros resultados que aparecen en Google es [esta página de github](https://markdowncss.github.io).

Yo he elegido el estilo _Modest_ cuyo css se encuentra en el siguiente [enlace](https://github.com/markdowncss/modest/blob/master/css/modest.css).

Una vez descargado el fichero `modest.css` lo colocamos en el mismo directorio que el fichero markdown que queremos convertir, por ejemplo este mismo documento `Markdown.md`.

Y ejecutamos el comando:

```bash
pandoc --css=modest.css -s -f markdown+smart --toc --metadata -pagetitle="Ejemplo markdown" --to=html5 ejemplo.md -o ejemplo.html
```

## Ejercicio

### Apartado 1

Usando el editor de código VSCode, crea un archivo con extensión `.md` y escribe un documento que incluya los elementos básicos mencionados anteriormente. Luego, utiliza la función de vista previa de VSCode para ver cómo se renderiza tu documento en HTML.

Comprueba que si incluyes una imagen esta puede estar asociada a un fichero local o a una URL externa.

### Apartado 2

Incluye ahora en tu documento una tabla y un bloque de código. Asegúrate de que ambos se renderizan correctamente en la vista previa de VSCode.

### Apartado 3

Convierte el documento a HTML utilizando una herramienta en línea o un programa de conversión de Markdown a HTML. Compara el resultado con la vista previa de VSCode para asegurarte de que todo se ha convertido correctamente.
