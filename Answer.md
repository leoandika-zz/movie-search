Question : 

```
Jika ada database table "USER" yg memiliki 3 kolom: ID, name, parent
Kolom ID adalah Primary Key
Kolom UserName adalah Nama User
Kolom Parent adalah ID dari User yang menjadi Creator untuk user tertentu.
eg. 
——————————————————————————
| ID | UserName | Parent |
——————————————————————————
| 1  | Ali      |   2    | 
| 2  | Budi     |   0    |
| 3  | Cecep    |   1    |
—————————————————————————-
Berarti record Ali dicreate oleh Budi, Cecep dicreate Ali, Sementara Budi tidak dicreate siapapun
Tuliskan SQL query untuk mendapatkan data berisi:
ID, UserName, ParentUserName
Kolom ParentUserName adalah UserName berdasarkan value Parent
```

Answer :
```sql
SELECT us.ID, us.UserName, 
CASE
	WHEN us2.UserName IS NULL THEN 'None'
	ELSE us2.UserName
END AS 'ParentUserName'
FROM User us 
LEFT JOIN User us2 ON us.Parent = us2.ID
ORDER BY us.ID
```
Asumsi : ketika sebuah record tidak di create siapapun, maka `ParentUserName` akan diisi sebagai `None`