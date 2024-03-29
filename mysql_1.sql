/* To create a new database: */
CREATE DATABASE mathi;

/* Ex : */
CREATE DATABASE sangeetha;

/* To drop(Remove) a database:*/
DROP DATABASE mathi;

/* Ex:*/
DROP DATABASE sangeetha;


/* To view all databases:*/
SHOW DATABASES;

/* To select a database:*/
USE mathi;

/*  Ex:*/
USE sangeetha;


/*  To create a new table:*/

CREATE TABLE users(id INT NOT NULL AUTO_INCREMENT, name VARCHAR(50) NOT NULL, age INT NOT NULL,PRIMARY KEY(id));

/*  To view all tables in a database:*/
SHOW TABLES;
/* To view structure of a table:*/
DESCRIBE users;


/* To add a new column inside a table:*/
ALTER TABLE users ADD gender VARCHAR(10) NOT NULL AFTER AGE;

/* To add multiple columns to a table:*/
ALTER TABLE users ADD city VARCHAR(50) NOT NULL, ADD contact VARCHAR(50) NOT NULL;

/*To modify a column in a table:*/
ALTER TABLE users MODIFY contact VARCHAR(25) NOT NULL;

/*To rename a table:*/
ALTER TABLE users RENAME TO students;

/*To view recoords in a table:*/
SELECT * FROM students;

/* To insert a record into a table:*/
INSERT INTO students VALUES(NULL,'Ram',25,'Male','Salem','9874563210');

/* To insert mutiple records into a table:*/
INSERT INTO students(name,age,gender,city,contact) VALUES ('Ravi',23,'Male','Namakkal','9876543210'),('Sara',23,'Female','Erode','9874521360');

/* To delete a row from table:*/
DELETE FROM students WHERE id=4;

/* To Modify/Update a row in table:*/
UPDATE students SET city='Hosur',contact='9988776655' WHERE id=3;

/* To clear all data from a table:*/
TRUNCATE TABLE students;

/* Select particular fields in a table:*/
SELECT name,age FROM students;

/* elect records with criteria:*/
SELECT name,age,city FROM students WHERE city='Hosur';

/* Select recoords with multiple criteria:*/
SELECT name,age,city FROM students WHERE city='Hosur' AND age >= 23;

SELECT name,age,city FROM students WHERE city='Namakkal' OR city='Hosur';

SELECT name,age,city FROM students WHERE (city='Namakkal' OR city='Hosur') AND age >= 23;

/* To select unique values in a field:*/
SELECT DISTINCT city FROM students;

/* To count unique values in a field:*/
SELECT COUNT(DISTINCT city) FROM students;

/* To count unique values in a field with allies name:*/

SELECT COUNT(DISTINCT city) AS total FROM students;

/* To Select N number of recoords:*/
SELECT * FROM students LIMIT 5;

/* To Select a range of recoords:*/
SELECT * FROM students LIMIT 5,15;

/* To select the very first record from a table:*/
SELECT * FROM students LIMIT 0,1;

/* To select the very last record from a table:*/
SELECT * FROM students ORDER BY id DESC LIMIT 0,1;

/* To select maximum value in a table:*/
SELECT MAX(age) FROM students;

/* To select minimum value in a table:*/
SELECT MIN(age) FROM students;

/* To select average of a field:*/
SELECT AVG(age) FROM students;

/* To round a float value:*/
SELECT ROUND(AVG(age),0) FROM students;

/* To select sum of a field:*/
SELECT SUM(age) FROM students;

/* To select group wise records:*/
SELECT gender,COUNT(id) FROM students GROUP BY gender;


/* WILD CARD QUERIES:
To select values starts with some letters:*/
SELECT NAME FROM students WHERE name LIKE 'Ra%';

/* To select values ends with some letters:*/
SELECT NAME FROM students WHERE name LIKE '%am';

/* To select values contains some letters:*/
SELECT NAME FROM students WHERE name LIKE '%la%';

/* To select values having given keywords:*/
SELECT * FROM STUDENTS WHERE city IN('Salem','Namakkal');

/* To select values not having given keywords:*/
SELECT * FROM STUDENTS WHERE city NOT IN('Salem','Namakkal');

/* To select values not starts with some letters:*/
SELECT NAME FROM students WHERE name NOT LIKE 'Ra%';

/* To select values using between keyword:*/
SELECT name,age FROM students WHERE age BETWEEN 24 AND 30;


/* JOIN QUERIES:
1.Inner Join(To select common values only)*/
SELECT emp.name,emp.design,salary.sdate,salray.amt FROM emp INNER JOIN salary ON emp.id = salary.id;

/* 2.Left Join(To select common values and all values from the first table)*/
SELECT emp.name,emp.design,salary.sdate,salray.amt FROM emp LEFT JOIN salary ON emp.id = salary.id;

/* 3.Right Join(To select common values and all values from the second table)*/
SELECT emp.name,emp.design,salary.sdate,salray.amt FROM emp LEFT JOIN salary ON emp.id = salary.id;

/* To select with cases:*/
SELECT name,city,(CASE WHEN city='Salem' THEN 100 WHEN city='Namakkal' THEN 200 WHEN city='Chennai' THEN 300 WHEN city='Hosur' THEN 400 ELSE 0 END) AS amt FROM students;

/*Multiple Inner Join:*/
SELECT students.name,students.rollno,course.cname,marks.M1,marks.M2,marks.M3 FROM students INNER JOIN course ON students.cid=course.cid INNER JOIN marks ON students.id=marks.id;




/*Simple Dynamic Marklist using Queries:*/

SELECT students.name,students.rollno,
course.cname as Course,
marks.M1,marks.M2,marks.M3,
(marks.M1,marks.M2,marks.M3) AS Total,
ROUND(((marks.M1,marks.M2,marks.M3)/3),2) AS Average
CASE
	WHEN marks.M1 >= 35 AND marks.M2 >= 35 AND marks.M3 >= 35 THEN 'PASS'
	ELSE 'FAIL'
END AS Result,
CASE
	WHEN marks.M1 >= 35 AND marks.M2 >= 35 AND marks.M3 >= 35 THEN
		CASE
			WHEN ROUND(((marks.M1,marks.M2,marks.M3)/3),2) >= 90 AND ROUND(((marks.M1,marks.M2,marks.M3)/3),2) < = 100 THEN 'A'
			WHEN ROUND(((marks.M1,marks.M2,marks.M3)/3),2) >= 80 AND ROUND(((marks.M1,marks.M2,marks.M3)/3),2) < = 89 THEN 'B'
			ELSE
			'C'
		END
	ELSE 'NO GRADE'
END AS Result,
 FROM students INNER JOIN course ON students.cid=course.cid INNER JOIN marks ON students.id=marks.id;
 
 
 
/*To use dynamically Created column names:*/

 SELECT students.name,students.rollno,
course.cname as Course,
marks.M1,marks.M2,marks.M3,
(marks.M1,marks.M2,marks.M3) AS Total,
ROUND(((marks.M1,marks.M2,marks.M3)/3),2) AS Average
CASE
	WHEN marks.M1 >= 35 AND marks.M2 >= 35 AND marks.M3 >= 35 THEN 'PASS'
	ELSE 'FAIL'
END AS Result,
CASE
	WHEN marks.M1 >= 35 AND marks.M2 >= 35 AND marks.M3 >= 35 THEN
		CASE
			WHEN ROUND(((marks.M1,marks.M2,marks.M3)/3),2) >= 90 AND ROUND(((marks.M1,marks.M2,marks.M3)/3),2) < = 100 THEN 'A'
			WHEN ROUND(((marks.M1,marks.M2,marks.M3)/3),2) >= 80 AND ROUND(((marks.M1,marks.M2,marks.M3)/3),2) < = 89 THEN 'B'
			ELSE
			'C'
		END
	ELSE 'NO GRADE'
END AS Result,
 FROM students INNER JOIN course ON students.cid=course.cid INNER JOIN marks ON students.id=marks.id WHERE course.cname='BCA' having Result = 'PASS' AND (Average >= 70 AND Average <= 100);
 
 
 
/*Creating queries as View:*/
 
CREATE VIEW Reports AS
SELECT students.name,students.rollno,
course.cname as Course,
marks.M1,marks.M2,marks.M3,
(marks.M1,marks.M2,marks.M3) AS Total,
ROUND(((marks.M1,marks.M2,marks.M3)/3),2) AS Average
CASE
	WHEN marks.M1 >= 35 AND marks.M2 >= 35 AND marks.M3 >= 35 THEN 'PASS'
	ELSE 'FAIL'
END AS Result,
CASE
	WHEN marks.M1 >= 35 AND marks.M2 >= 35 AND marks.M3 >= 35 THEN
		CASE
			WHEN ROUND(((marks.M1,marks.M2,marks.M3)/3),2) >= 90 AND ROUND(((marks.M1,marks.M2,marks.M3)/3),2) < = 100 THEN 'A'
			WHEN ROUND(((marks.M1,marks.M2,marks.M3)/3),2) >= 80 AND ROUND(((marks.M1,marks.M2,marks.M3)/3),2) < = 89 THEN 'B'
			ELSE
			'C'
		END
	ELSE 'NO GRADE'
END AS Result,
FROM students INNER JOIN course ON students.cid=course.cid INNER JOIN marks ON students.id=marks.id;
 
/*Accessing views like a table:*/
 SELECT * FROM Reports;
 SELECT * FROM Reports WHERE Result = 'PASS';
 
 
 /*Update table using Inner Join:*/
 
 UPDATE marks INNER JOIN students ON students.id = marks.id SET M1 = 100,M2 = 100,M3 = 100 WHERE students.rollno = 'A1001';
 
 
/*Creating Triggers:*/
CREATE TRIGGER before_products_update
	BEFORE UPDATE ON product
	FOR EACH ROW
BEGIN 
	INSERT INTO price_logs(pid,price,new_price)
	VALUES(old.pid,old.rate,new.rate);
END$$
