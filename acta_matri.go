// pkg/acta/acta.go
package acta_matri.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    DatosRegistrador    uint      `json:"user_id,omitempty"`
    LugarHrFecha    uint      `json:"user_id,omitempty"`
    DatosConyugeM    uint      `json:"user_id,omitempty"`
    DatosConyugeF    uint      `json:"user_id,omitempty"`
    DDConyuges    uint      `json:"user_id,omitempty"`
    AceptacionExpresa    uint      `json:"user_id,omitempty"`
    CapitulacionesMatrimoniales    uint      `json:"user_id,omitempty"`
    DatosApoderado    uint      `json:"user_id,omitempty"`
    DatosHijos    uint      `json:"user_id,omitempty"`
    DatosTestigo    uint      `json:"user_id,omitempty"`
    DatosActa    uint      `json:"user_id,omitempty"`
    CircunstanciasEspe    uint      `json:"user_id,omitempty"`
    NotaMarginal    uint      `json:"user_id,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Registrador.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Oficina    string      `json:"Oficina,omitempty"`
    Resolucion    uint      `json:"Resolucion,omitempty"`
    Fecha    string      `json:"Fecha,omitempty"`
    Resolucion    uint      `json:"Resolucion,omitempty"`
    Fecha    string      `json:"Fecha,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Lugar_Hr_Fecha.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Lugar    string      `json:"Nombres,omitempty"`
    Hora    string      `json:"Apellidos,omitempty"`
    Fecha    uint      `json:"DocuIde,omitempty"`
    DatosExpediente    string      `json:"Oficina,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_ConyugeM.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    FechaNaci    string      `json:"FechaNaci,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    EdoNaci    string      `json:"PaisNaci,omitempty"`
    MnpioNaci    string      `json:"Edad,omitempty"`
    ParrNaci    uint      `json:"Nacionalidad,omitempty"`
    Nacionalidad    string      `json:"PaisNaci,omitempty"`
    EdoResid    string      `json:"PaisNaci,omitempty"`
    MnpioResid    string      `json:"Edad,omitempty"`
    ParrResid    uint      `json:"Nacionalidad,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    NombresApellidosPadre    string      `json:"Nombres,omitempty"`
    DocuIdePadre    uint      `json:"DocuIde,omitempty"`
    ProfesionPadre    string      `json:"Profesion,omitempty"`
    ResidenciaPadre    uint      `json:"Comunidad,omitempty"`
    NombresApellidosMadre    string      `json:"Nombres,omitempty"`
    DocuIdeMadre    uint      `json:"DocuIde,omitempty"`
    ProfesionMadre    string      `json:"Profesion,omitempty"`
    ResidenciaMadre    uint      `json:"Comunidad,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_ConyugeF.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    FechaNaci    string      `json:"FechaNaci,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    EdoNaci    string      `json:"PaisNaci,omitempty"`
    MnpioNaci    string      `json:"Edad,omitempty"`
    ParrNaci    uint      `json:"Nacionalidad,omitempty"`
    Nacionalidad    string      `json:"PaisNaci,omitempty"`
    EdoResid    string      `json:"PaisNaci,omitempty"`
    MnpioResid    string      `json:"Edad,omitempty"`
    ParrResid    uint      `json:"Nacionalidad,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    NombresApellidosPadre    string      `json:"Nombres,omitempty"`
    DocuIdePadre    uint      `json:"DocuIde,omitempty"`
    ProfesionPadre    string      `json:"Profesion,omitempty"`
    ResidenciaPadre    uint      `json:"Comunidad,omitempty"`
    NombresApellidosMadre    string      `json:"Nombres,omitempty"`
    DocuIdeMadre    uint      `json:"DocuIde,omitempty"`
    ProfesionMadre    string      `json:"Profesion,omitempty"`
    ResidenciaMadre    uint      `json:"Comunidad,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package DD_Conyuges.go

import "time"

// Post created by a default text.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`


    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Aceptacion_Expresa.go

import "time"

// Post created by a default text.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`


    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Capitulaciones_Matrimoniales.go

import "time"

// Post created by a default text.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    DatosRegistro    string      `json:"Nombres,omitempty"`
    Nro    string      `json:"Nombres,omitempty"`
    Fecha    string      `json:"Nombres,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Apoderado.go

import "time"

// Post created by a default text.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    DatosNotaria    uint      `json:"DocuIde,omitempty"`
    Nro    string      `json:"Nombres,omitempty"`
    Fecha    string      `json:"Nombres,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

    // pkg/acta/acta.go
package Datos_Hijos.go

import "time"

// Post created by a default text.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    NombresApellidoshijo1    string      `json:"Direccion,omitempty"`
    DocuIdehijo1    uint      `json:"DocuIde,omitempty"`
    EdadHijo1    string      `json:"Apellidos,omitempty"`
    ActaNaciNro1    string      `json:"Apellidos,omitempty"`
    FechaNaci1    string      `json:"Apellidos,omitempty"`
    NombresApellidoshijo2    string      `json:"Direccion,omitempty"`
    DocuIdehijo2    uint      `json:"DocuIde,omitempty"`
    EdadHijo2    string      `json:"Apellidos,omitempty"`
    ActaNaciNro2    string      `json:"Apellidos,omitempty"`
    FechaNaci2    string      `json:"Apellidos,omitempty"`
    NombresApellidoshijo3    string      `json:"Direccion,omitempty"`
    DocuIdehijo3    uint      `json:"DocuIde,omitempty"`
    EdadHijo3    string      `json:"Apellidos,omitempty"`
    ActaNaciNro3    string      `json:"Apellidos,omitempty"`
    FechaNaci3    string      `json:"Apellidos,omitempty"`
    NombresApellidoshijo4    string      `json:"Direccion,omitempty"`
    DocuIdehijo4    uint      `json:"DocuIde,omitempty"`
    EdadHijo4    string      `json:"Apellidos,omitempty"`
    ActaNaciNro4    string      `json:"Apellidos,omitempty"`
    FechaNaci4    string      `json:"Apellidos,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Testigo.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Edad    string      `json:"Edad,omitempty"`
    Nacionalidad    uint      `json:"Nacionalidad,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    Nombres2    string      `json:"Nombres,omitempty"`
    Apellidos2    string      `json:"Apellidos,omitempty"`
    DocuIde2    uint      `json:"DocuIde,omitempty"`
    Profesion2    string      `json:"Profesion,omitempty"`
    Edad2    string      `json:"Edad,omitempty"`
    Nacionalidad2    uint      `json:"Nacionalidad,omitempty"`
    Direccion2    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Acta.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Numero    string      `json:"Numero,omitempty"`
    Fecha    string      `json:"Fecha,omitempty"`
    Autoridad    uint      `json:"Autoridad,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Circunstancias_Espe.go

import "time"

// Post created by a note.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nota    string      `json:"InfoConceProte,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

    // pkg/acta/acta.go
package Nota_Marginal.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nota    string      `json:"InfoConceProte,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;
    