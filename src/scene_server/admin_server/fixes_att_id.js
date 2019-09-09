function getNextSequence(name) {
    var ret = db.cc_idgenerator.findAndModify(
        {
            query: { _id: name },
            update: { $inc: { SequenceID: 1 } },
            new: true
        }
    );
    return ret.SequenceID;
}

db.cc_ObjAttDes.find().forEach(function (myDoc) {
    var nid = getNextSequence("cc_ObjAttDes")
    print("nid: " + nid);
    db.cc_ObjAttDes.update(
        {
            "obj_id": myDoc.obj_id,
            "bk_property_id": myDoc.bk_property_id,
            "org_id": myDoc.org_id,
        },
        {
            "$set": {
                "id": nid,
            }
        })
});

