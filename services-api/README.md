# Sistema de Gestión de Reservas para Servicios de Mantenimiento

**Características**:

- Registro y autenticación de usuarios
- CRUD de reservas, talleres y servicios disponibles
- Implementación de una lógica para verificar la disponibilidad de horarios

## Patrón de Diseño

Para este proyecto, utilizaremos una combinación de los patrones MVC (Model-View-Controller) y Repository para separar las responsabilidades de manera clara:

- Models: Representan las estructuras de datos que mapean a la base de datos.
- Views: En este caso, como estamos creando una API, no manejaremos vistas tradicionales, sino que las respuestas JSON actuarán como nuestras "vistas".
- Controllers: Manejan las solicitudes HTTP y coordinan las respuestas utilizando los modelos y servicios.
- Repositories: Proporcionan una abstracción sobre las operaciones de la base de datos.
- Services: Encapsulan la lógica de negocio y se comunican con los repositories para obtener o modificar datos.
