from tarantool import Connection
c = Connection(
    "89.208.85.244", 
    31320,
    user='admin', 
    password='34a7d3bf316730ae'
)
result = c.insert("customer", (332, 'John Smith'))
space = c.space("customer")
results = space.select()
print(results)

""" curl -X POST --data "fullname= Taran Tool" https://dnbrekfrkkm.try.tarantool.io/add_user
{"status":"Success!","result":{"rows":[[10568,"d23f870c-ef1c-436d-942a-5266758f90e0"," Taran Tool"]],"metadata":[{"type":"unsigned","name":"bucket_id","is_nullable":false},{"type":"uuid","name":"user_id","is_nullable":false},{"type":"string","name":"fullname","is_nullable":false}]}} """

""" curl -X POST --data "description=My first tiktok" https://dnbrekfrkkm.try.tarantool.io/add_video
{"status":"Success!","result":{"rows":[[12121,"8f763048-e21f-4966-9036-ab37a455721d","My first tiktok"]],"metadata":[{"type":"unsigned","name":"bucket_id","is_nullable":false},{"type":"uuid","name":"video_id","is_nullable":false},{"type":"string","name":"description","is_nullable":true}]}}r 

After app

"""