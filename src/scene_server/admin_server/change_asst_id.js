// params:
var obj_asst_id = "a_default_b"
var newAsstID = "run"
var org_id = "0"

// execute
var objAsst = db.cc_ObjAsst.findOne({ "obj_asst_id": obj_asst_id, "org_id": org_id })
if (objAsst != null) {

    var new_obj_asst_id = objAsst.obj_id + "_" + newAsstID + "_" + objAsst.asst_obj_id
    db.cc_ObjAsst.update({ "obj_asst_id": obj_asst_id, "org_id": org_id }, { "$set": { "asst_id": newAsstID, "obj_asst_id": new_obj_asst_id } })

    db.cc_InstAsst.update({ "obj_asst_id": obj_asst_id, "org_id": org_id }, { "$set": { "asst_id": newAsstID, "obj_asst_id": new_obj_asst_id } }, { "multi": 1 })
}
