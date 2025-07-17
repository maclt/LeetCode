-- Write your PostgreSQL query statement below

SELECT name
FROM Employee
JOIN (
    SELECT managerId, COUNT(managerId)
    FROM Employee
    GROUP BY managerId
    HAVING COUNT(managerId) >= 5
) as AtLeastFiveDirects
 on Employee.id = AtLeastFiveDirects.managerId
