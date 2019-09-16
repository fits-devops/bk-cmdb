db.cc_ObjAsst.find().forEach(function (objasst) {
    var ret = db.cc_InstAsst.update(
        {
            "obj_id": objasst.obj_id,
            "asst_obj_id": objasst.asst_obj_id,
            "org_id": objasst.org_id
        },
        {
            "$set": {
                "obj_asst_id": objasst.obj_asst_id,
                "asst_id": objasst.asst_id,
                "last_time": new Date(),
            }
        })
    print(ret)
});
