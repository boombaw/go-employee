package query

// Query untuk menyisipkan data karyawan
var InsertEmployee = `INSERT INTO employee (
                        name,
                        department,
                        status,
                        salary
                      )
                      VALUES (?,?,?,?)`

// Query untuk memperbarui data karyawan
var UpdateEmployee = `UPDATE employee
                      SET
                        name = ?,
                        department = ?,
                        status = ?,
                        salary = ?
                      WHERE
                        id = ?`

// Query untuk memilih data karyawan berdasarkan ID
var SelectEmpByID = `SELECT *
                     FROM
                       employee
                     WHERE
                       id = ?`

// Query untuk memilih data karyawan berdasarkan status
var SelectEmpByStatus = `SELECT *
                         FROM
                           employee
                         WHERE
                           status = ?`

// Query untuk memilih semua data karyawan
var SelectAllEmp = `SELECT *
                    FROM
                      employee
                    LIMIT ? OFFSET ?`

// Query untuk menghitung total karyawan
var TotalEmp = `SELECT 
					count(*)
                FROM
                  employee`

// Query untuk memeriksa duplikat nama di departemen
var CheckDuplicateName = `SELECT COUNT(1)
                          FROM
                            employee
                          WHERE
                            name = ?
                            AND department = ?`

// Query untuk mendapatkan rata-rata gaji berdasarkan departemen
var AvgSalary = `SELECT
                   AVG(salary) AS salary,
                   department
                FROM
                	employee
				WHERE department = ?
                 GROUP BY
                   department;`
