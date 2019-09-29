db.cc_ObjAsst.find({ "asst_id": "mainline", "obj_id": {"$nin": ["set","module","host"]} }).forEach(function (myDoc) {
    var ret = db.cc_ObjAsst.remove({"asst_id": "mainline", "obj_id": myDoc.obj_id})
    print("delete cc_ObjAsst for ", myDoc.obj_id, " result: ", ret)
    ret = db.cc_ObjAttDes.remove({"obj_id": myDoc.obj_id})
    print("delete cc_ObjAttDes for ", myDoc.obj_id, " result: ", ret)
    ret = db.cc_ObjectBase.remove({"obj_id": myDoc.obj_id})
    print("delete cc_ObjectBase for ", myDoc.obj_id, " result: ", ret)
});

var ret = db.cc_ObjAsst.update({"asst_id": "mainline", "obj_id": "set"}, {"$set": {"asst_obj_id": "biz"}});
print("update cc_ObjAsst for ", "set", " result: ", ret)

db.cc_SetBase.find().forEach(function (myDoc) {
    ret = db.cc_SetBase.update({"set_id": myDoc.set_id,"org_id": myDoc.org_id}, {"$set": {"bk_parent_id": myDoc.biz_id}})
    print("update cc_SetBase for ", myDoc.set_id, " result: ", ret)
});

print("done")
