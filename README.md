# TEST_GOLANG_NODEJS_DOCKER
JX TEAM TEST

LINK POSTMAN:
https://www.getpostman.com/collections/aa928ee4d36a78299286

FLOW GOLANG SOURCE:
1. RUN GOLANG - CRAWLER DATA FROM FLIGHTRADAR24 and then save them to MongoDB
2. CALL API localhost:8080/api/get-all-crawler-data (to SHOW DATA UNPROCESS From MongoDB)
3. CALL API localhost:8080/api/process-data-to-postgres (to GET DATA UNPROCESS From MongoDB and PROCESS THEM and then SAVE THEM TO POSTGRES)

NODEJS SOURCE has 2 API:
1. GET ALL DATA FLIGHT
2. GET DATA TRACK FROM ID FLIGHT (LIST TRACK là danh sách các điểm tọa độ mà chuyến bay đó đi qua)

