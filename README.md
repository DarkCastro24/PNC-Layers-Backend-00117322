# Programacion N Capas - Backend API (Tarea)

**Estudiante:** Diego Eduardo Castro Quintanilla  
**Carnet:** 00117322  

---

Lista de endpoints

---

### Obtener todos los usuarios

- **Metodo:** GET  
- **URL:** http://localhost:8080/users

### Obtener usuario por ID

- **Metodo:** GET  
- **URL:** http://localhost:8080/users/{id}

### Crear un nuevo usuario 

- **Metodo:** POST 
- **URL:** http://localhost:8080/users
- **Body: {
  "name": "Diego",
  "email": "diego@gmail.com"
}**

### Actualizar un usuario 

- **Metodo:** PUT
- **URL:** http://localhost:8080/users/{id}
- **Body: {
  "name": "Diego Actualizado",
  "email": "diego2@gmail.com"
}**

### Eliminar un usuario 

- **Metodo:** DELETE
- **URL:** http://localhost:8080/users/{id}

### PD: Trate de hacer el JWT pero no logre implementarlo :(
