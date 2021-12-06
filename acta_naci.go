// pkg/acta/acta.go
package acta_naci.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    DatosRegistrador    uint      `json:"user_id,omitempty"`
    DatosPresentado    uint      `json:"user_id,omitempty"`
    DatosCertificado    uint      `json:"user_id,omitempty"`
    DatosMadre    uint      `json:"user_id,omitempty"`
    DatosPadre    uint      `json:"user_id,omitempty"`
    DatosDeclarante    uint      `json:"user_id,omitempty"`
    DatosTestigos    uint      `json:"user_id,omitempty"`
    DatosActa    uint      `json:"user_id,omitempty"`
    InscripcionProteccion    uint      `json:"user_id,omitempty"`
    InscripcionJudicial    uint      `json:"user_id,omitempty"`
    InscripcionExtem    uint      `json:"user_id,omitempty"`
    CircunstanciasEspe    uint      `json:"user_id,omitempty"`
    DocuPresentados    uint      `json:"user_id,omitempty"`
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
package Datos_Presentado.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    FechaNaci    string      `json:"FechaNaci,omitempty"`
    Sexo    string      `json:"Sexo,omitempty"`
    HoraNaci    string      `json:"HoraNaci,omitempty"`
    PaisNaci    string      `json:"PaisNaci,omitempty"`
    EdoNaci    string      `json:"EdoNaci,omitempty"`
    muniNaci    string      `json:"muniNaci,omitempty"`
    ParroquiaNaci    string      `json:"ParroquiaNaci,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Certificado.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Certficado    string      `json:"Certficado,omitempty"`
    FechaExpedicion    string      `json:"FechaExpedicion,omitempty"`
    Autoridad    string      `json:"Autoridad,omitempty"`
    Sexo    string      `json:"Sexo,omitempty"`
    NMPPS    string      `json:"NMPPS,omitempty"`
    CentroSalud    string      `json:"CentroSalud,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Madre.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Edad    string      `json:"Edad,omitempty"`
    Nacionalidad    uint      `json:"Nacionalidad,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Comunidad    uint      `json:"Comunidad,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Padre.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Edad    string      `json:"Edad,omitempty"`
    Nacionalidad    uint      `json:"Nacionalidad,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Comunidad    uint      `json:"Comunidad,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Declarante.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Edad    string      `json:"Edad,omitempty"`
    Nacionalidad    uint      `json:"Nacionalidad,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Comunidad    uint      `json:"Comunidad,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Datos_Testigos.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Nombres    string      `json:"Nombres,omitempty"`
    Apellidos    string      `json:"Apellidos,omitempty"`
    DocuIde    uint      `json:"DocuIde,omitempty"`
    Edad    string      `json:"Edad,omitempty"`
    Nacionalidad    uint      `json:"Nacionalidad,omitempty"`
    Profesion    string      `json:"Profesion,omitempty"`
    Comunidad    uint      `json:"Comunidad,omitempty"`
    Direccion    string      `json:"Direccion,omitempty"`
    Nombres2    string      `json:"Nombres,omitempty"`
    Apellidos2    string      `json:"Apellidos,omitempty"`
    DocuIde2    uint      `json:"DocuIde,omitempty"`
    Edad2    string      `json:"Edad,omitempty"`
    Nacionalidad2    uint      `json:"Nacionalidad,omitempty"`
    Profesion2    string      `json:"Profesion,omitempty"`
    Comunidad2    uint      `json:"Comunidad,omitempty"`
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
package Inscripcion_Proteccion.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    ConcejoProteNiños    string      `json:"ConcejoProteNiños,omitempty"`
    Medida    string      `json:"Medida,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    NombresApellidosConce    uint      `json:"NombresApellidosConce,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Inscripcion_Judicial.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    Tribunal    string      `json:"Tribunal,omitempty"`
    SentenciaN    string      `json:"SentenciaN,omitempty"`
    NombresApellidosJuez    uint      `json:"NombresApellidosJuez,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`;

// pkg/acta/acta.go
package Inscripcion_Extem.go

import "time"

// Post created by a user.
type Post struct {
    ID        uint      `json:"id,omitempty"`
    Body      string    `json:"body,omitempty"`
    InfoConceProte    string      `json:"InfoConceProte,omitempty"`
    Autoridad    string      `json:"Autoridad,omitempty"`
    Fecha    uint      `json:"Fecha,omitempty"`
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
package Docu_Presentados.go

import "time"

// Post created by a user.
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
    