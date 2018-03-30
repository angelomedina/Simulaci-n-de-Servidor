/* Created by "go tool cgo" - DO NOT EDIT. */

/* package command-line-arguments */


#line 1 "cgo-builtin-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

typedef struct { const char *p; ptrdiff_t n; } _GoString_;

#endif

/* Start of preamble from import "C" comments.  */




/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

typedef _GoString_ GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


//agrega autos a sus rutas

extern void cargaPaquetesAutos(GoString p0, GoInt p1);

//carga las rutas con sus eventuales destinos

extern void caragarRutas();

//genera el random del destino de la ruta origen

extern GoInt random(GoInt p0, GoInt p1);

//funcion para igualar strings

extern GoInt equal(GoString p0, GoString p1);

//simula la carga vehiclar el la rotonda

extern void servidor();

//retorna la cantidad de autos en cola  de cad ruta

extern GoInt lenghtCola(GoSlice p0);

//ejecula la cola delas rutas/solicitud de acceso

extern void solicitudCola(GoSlice p0, GoInt p1);

//ejecuta el acceso de la solicitud

extern void respuestaCola(GoSlice p0);

#ifdef __cplusplus
}
#endif
