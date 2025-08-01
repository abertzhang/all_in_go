```
// HMYConfirmEntryInfoController类
 confirmEntryInfo(pass) async {
    if (ObjectUtil.isEmpty(actualEntryTime.value)) {
      HMYEasyLoading.showToast('请选择预约进场时间');
      return;
    }
    if (ObjectUtil.isEmpty(confirmDeliverDeviceAddressDetail?.addressDetail)) {
      HMYEasyLoading.showToast('交车地址不完整\n请重新选择');
      return;
    }
    String reason = '';
```

```
   // HMYEditEnterAreaController类
    if (emptyDevice) {
      HMYEasyLoading.showToast('进场设备不能为空');
      return false;
    }

    if (ObjectUtil.isEmpty(deliverDeviceAddressDetail?.addressDetail)) {
      HMYEasyLoading.showToast('交车地址不完整\n请重新选择');
      return false;
    }

    if (submitModel.value.customerEntryTime == null) {
      HMYEasyLoading.showToast('进场时间不能为空');
      return false;
    }
```

```
    //HMYInitEnterAreaController
    if (emptyDevice) {
      HMYEasyLoading.showToast('进场设备不能为空');
      return false;
    }
    if (ObjectUtil.isEmpty(deliverDeviceAddressDetail?.longitude)) {
      HMYEasyLoading.showToast('交车地址不完整,\n请重新选择');
      return false;
    }

    if (submitModel.value.customerEntryTime == null) {
      HMYEasyLoading.showToast('进场时间不能为空');
```

```
//HMYCreateExitController
 //监听地图地址
    subscription = HMYEventUtil.to.on<HMYAMapAddressBean>().listen((addressBean) {
      if (addressBean.typeId != idExitAddressDetail) return;
      detailModel.exitRequirementInfo?.exitAddressDetail = addressBean;
      exitAddress = addressBean.mapAddress ?? '';
      update([idExitAddressDetail]);
    });
  
  **************************
      if (!hasExitCount) {
      HMYEasyLoading.showToast('请增加退场数量');
      return;
    }
    if (ObjectUtil.isEmpty(detailModel.exitRequirementInfo?.exitAddressDetail?.addressDetail)) {
      HMYEasyLoading.showToast('退场地址不完整,\n请重新选择');
      return;
    }
    if (exitAddress.isEmpty) {
      HMYEasyLoading.showToast('请输入退场地址');
      return;
    }
```

```
//HMYExitConfirmInfoController
bool _canPush() {
    if (ObjectUtil.isEmpty(actualExitAddressDetail?.addressDetail)) {
      HMYEasyLoading.showToast('退场地址不完整,\n请重新选择');
      return false;
    }
    if (confirmAddress.isEmpty) {
      HMYEasyLoading.showToast('请输入退场地址');
      return false;
    }
    if (confirmTime.isEmpty) {
      HMYEasyLoading.showToast('请选择退场时间');
      return false;
    }
    return true;
  }
```

```
     Icon(Icons.photo_camera_outlined, size: width / 3, color: const Color(0xff999999)),
            Text(
              countSelected == 0 ? '上传图片\n(最多${widget.maxCount}张)':'$countSelected/${widget.maxCount}',
              style:  TextStyle(fontSize: 10.sp, color: Color(0xFFA2AABD)),
            ),
            
    -------        
   Widget build(BuildContext context) {
    return Wrap(
      runSpacing: 8,
      spacing: 8,
      runAlignment: WrapAlignment.spaceEvenly,           
```

