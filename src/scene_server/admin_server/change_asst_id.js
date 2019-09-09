// params:
var bk_obj_asst_id = "a_default_b"
var newAsstID = "run"
var org_id = "0"

// execute
var objAsst = db.cc_ObjAsst.findOne({ "bk_obj_asst_id": bk_obj_asst_id, "org_id": org_id })
if (objAsst != null) {

    var new_bk_obj_asst_id = objAsst.obj_id + "_" + newAsstID + "_" + objAsst.bk_asst_obj_id
    db.cc_ObjAsst.update({ "bk_obj_asst_id": bk_obj_asst_id, "org_id": org_id }, { "$set": { "bk_asst_id": newAsstID, "bk_obj_asst_id": new_bk_obj_asst_id } })

    db.cc_InstAsst.update({ "bk_obj_asst_id": bk_obj_asst_id, "org_id": org_id }, { "$set": { "bk_asst_id": newAsstID, "bk_obj_asst_id": new_bk_obj_asst_id } }, { "multi": 1 })
}
