# Sistema de Gestión de Reservas para Servicios de Mantenimiento

**Características**:

- Registro y autenticación de usuarios
- CRUD de reservas, talleres y servicios disponibles
- Implementación de una lógica para verificar la disponibilidad de horarios

## Patrón de Diseño

MVC (Model-View-Controller) y Repository:

- Models: Representan las estructuras de datos que mapean a la base de datos.
- Controllers: Manejan las solicitudes HTTP y coordinan las respuestas utilizando los modelos y servicios.
- Repositories: Proporcionan una abstracción sobre las operaciones de la base de datos.
- Services: Encapsulan la lógica de negocio y se comunican con los repositories para obtener o modificar datos.
- DI: Inyeccion de dependencias para manejar cada parte del sistema por separado.
