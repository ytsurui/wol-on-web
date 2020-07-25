
function machineMenuBuilder(machineData) {
    var base = $('<a class="dropdown-toggle" href="#" data-toggle="dropdown" aria-haspopup="true" area-expanded="true" />').text(machineData.name);
    var menu = $('<div class="dropdown-menu" />');
    
    var menuText = $('<h6 />').text(machineData.name);
    menu.append($('<div class="dropdown-item" />').append(menuText));
    menu.append($('<div class="dropdown-divider" />'));
    menu.append($('<div class="dropdown-item" onclick="sendWakeup(' + "'" + machineData.MacAddr + "', '" + machineData.netaddr + '\');" />').text("起動する"));
    menu.append($('<div class="dropdown-item" onclick="openConfigModal(' + "'" +  machineData.id + "'" + ');" />').text("設定する"));
    base.append(menu);

    return (base);
}

function getMachineList() {
    $("#machine-table").text("");
    $.ajax({
        type: "GET",
        url: "/api/machines"
    }).done(function(data, dataType) {

        //var machines = JSON.parse(data)
        machines = data;

        $.each(machines, function(i, machineData){
            var tr = $("<tr />");
            var mName = $("<td />").append(machineMenuBuilder(machineData));
            var mState = $("<td />");
            var mMac = $("<td />").text(machineData.MacAddr);
            var mIpAddr = $("<td />").text(machineData.ipaddr);
            var mNet = $("<td />").text(machineData.netaddr);

            tr.append(mName);
            tr.append(mState);
            tr.append(mMac);
            tr.append(mIpAddr);
            tr.append(mNet);

            $("#machine-table").append(tr);

            checkState(machineData.ipaddr, mState)
        })

    });
}

function checkState(ipaddr, respelem) {
    var mState = $("<span />").text("確認中");
    mState.addClass("badge badge-pill badge-light");
    respelem.empty();
    respelem.append(mState);

    $.ajax({
        type: "GET",
        url: "/api/ping?" + $.param({"ipaddr": ipaddr}),
        error: function (xhr, ajaxOptions, thrownError) {
            respelem.empty();
            if (xhr.status == 404) {
                var mState = $("<span />").text("停止中");
                mState.addClass("badge badge-pill badge-danger");
                respelem.append(mState);
            } else if (xhr.status != 200) {
                var mState = $("<span />").text("起動中");
                mState.addClass("badge badge-pill badge-success");
                respelem.append(mState);
            }
            checkStateInterval(ipaddr, respelem);
        }
    }).done(function(recvData) {
        respelem.empty();
        var mState = $("<span />").text("起動中");
        mState.addClass("badge badge-pill badge-success");
        respelem.append(mState);
        checkStateInterval(ipaddr, respelem);
    })
}

function checkStateInterval(ipaddr, respelem) {
    respelem.delay(60000).queue(function() {
        respelem.dequeue();
        checkState(ipaddr, respelem);
    })
}

function isString(obj) {
    return typeof (obj) == "string" || obj instanceof String;
};

function sendWakeup(macaddr, broadcastAddr) {
    var paramArray = {};
    paramArray.macaddr = macaddr;
    if (isString(broadcastAddr)) {
        if (broadcastAddr.length != 0) paramArray.broadcast = broadcastAddr;
    }

    $.ajax({
        type: "POST",
        url: "/api/wol?" + $.param(paramArray),
        error: function (xhr, ajaxOptions, thrownError) {
            if ((xhr.status == 404) || (xhr.status == 400) || (xhr.status == 500)) {
                alert("Wake on LANパケットの送信に失敗しました");
            }
        }
    }).done(function(recvData) {
        alert("Wake on LANパケットを送信しました");
    })
    
};

function openConfigModal(id) {
    $('#modalconfig').modal('show');
    var paramArray = {};
    paramArray.id = id;
    $.ajax({
        type: "GET",
        url: "/api/machines/item?" + $.param(paramArray),
        error: function (xhr, ajaxOptions, thrownError) {
            alert("設定の読み出しに失敗しました");
            $('#modalconfig').modal('hide');
        }
    }).done(function(recvData) {
        //var machineInfo = JSON.parse(recvData);
        machineInfo = recvData;
        configID = machineInfo.id;
        $("#edit-name").val(machineInfo.name);
        $("#edit-mac").val(machineInfo.MacAddr);
        $("#edit-ip").val(machineInfo.ipaddr);
        $("#edit-subnet").val(machineInfo.NetMask);
        $("#edit-wolnet").val(machineInfo.netaddr);

        $("#edit-name").removeClass("is-valid is-invalid");
        $("#edit-mac").removeClass("is-valid is-invalid");
        $("#edit-ip").removeClass("is-valid is-invalid");
        $("#edit-subnet").removeClass("is-valid is-invalid");

        showNetmask(machineInfo.NetMask);
        $("#edit-delbutton").show();
    })
}

function openNewModal() {
    $('#modalconfig').modal('show');
    configID = 0;

    $("#edit-name").val("");
    $("#edit-mac").val("");
    $("#edit-subnet").val("");
    $("#edit-ip").val("");
    $("#edit-wolnet").val("");

    $("#edit-name").removeClass("is-valid");
    $("#edit-name").addClass("is-invalid");
    $("#edit-mac").removeClass("is-valid");
    $("#edit-mac").addClass("is-invalid");
    $("#edit-ip").removeClass("is-valid");
    $("#edit-ip").addClass("is-invalid");
    $("#edit-subnet").removeClass("is-valid");
    $("#edit-subnet").addClass("is-invalid");

    $("#edit-delbutton").hide();

    showNetmask("");
}

function saveConfigModal() {
    var postData = {};
    postData.name = $("#edit-name").val();
    postData.MacAddr = $("#edit-mac").val();
    postData.ipaddr = $("#edit-ip").val();
    postData.NetMask = Number($("#edit-subnet").val());
    postData.netaddr = $("#edit-wolnet").val();

    if ((postData.name.length == 0) || (postData.MacAddr.length == 0)) {
        alert("必要な項目が入力されていません");
        return;
    }

    var paramArray = {};
    if (configID != 0) {
        paramArray.id = configID;
    }

    $.ajax({
        type: "POST",
        url: "/api/machines/item?" + $.param(paramArray),
        contentType: 'application/json',
        dataType: 'json',
        data: JSON.stringify(postData),
        error: function(xhr, ajaxOptions, thrownError) {
            if (xhr.status == 200) {
                // Close Modal
                $('body').removeClass('modal-open');
                $('.modal-backdrop').remove();
                $('#modalconfig').modal('hide');
                getMachineList();
            } else {
                alert("データの送信に失敗しました");
            }
        }
    }).done(function(recvData) {
        console.log(recvData);
    });
}

function deleteConfigModal() {
    if (configID == 0) return;
    var delconfirm = window.confirm("このマシンを削除しますか？");
    if (delconfirm) {
        $.ajax({
            type: "DELETE",
            url: "/api/machines/item?" + $.param({"id": configID}),
            contentType: "application/json",
            error: function(xhr, ajaxOptions, thrownError) {
                if (xhr.status == 200) {
                    // Close Modal
                    $('body').removeClass('modal-open');
                    $('.modal-backdrop').remove();
                    $('#modalconfig').modal('hide');
                    getMachineList();
                } else {
                    alert ("削除に失敗しました");
                }
            }
        }).done(function(recvData) {
            // Close Modal
            $('body').removeClass('modal-open');
            $('.modal-backdrop').remove();
            $('#modalconfig').modal('hide');
            getMachineList();
        })
    }
}

function editComputerName(e) {
    if (e.target.value.length == 0) {
        $("#edit-name").removeClass("is-valid");
        $("#edit-name").addClass("is-invalid");
    } else {
        $("#edit-name").removeClass("is-invalid");
        $("#edit-name").addClass("is-valid");
    }
}

function editMAC(e) {
    if (e.target.value.length == 0) {
        $("#edit-mac").removeClass("is-valid");
        $("#edit-mac").addClass("is-invalid");
    } else {
        $("#edit-mac").removeClass("is-invalid");
        $("#edit-mac").addClass("is-valid");
    }

    var r = /([a-f0-9]{2})([a-f0-9]{2})/i,
        str = e.target.value.replace(/[^a-f0-9]/ig, "");

    while (r.test(str)) {
        str = str.replace(r, '$1' + '-' + '$2');
    }

    e.target.value = str.slice(0, 17);
};


function formatIPaddr(e) {
    formatIP(e);
    if (e.target.value.length == 0) {
        $("#edit-subnet").val("");
        $("#edit-wolnet").val("");

        $("#edit-ip").removeClass("is-valid");
        $("#edit-ip").addClass("is-invalid");
        $("#edit-subnet").removeClass("is-valid");
        $("#edit-subnet").addClass("is-invalid");
        return;
    }

    $("#edit-ip").removeClass("is-invalid");
    $("#edit-ip").addClass("is-valid");

    var chkDot = e.target.value.indexOf('.');
    if (chkDot == -1) return;

    var ipelems = e.target.value.split('.');

    if (ipelems[0] == "") return;

    subnetValue = $("#edit-subnet").val();
    if (subnetValue.length == 0) {
        $("#edit-subnet").removeClass("is-valid is-invalid");
        if (ipelems[0] < 128) {
            // class A
            subnetValue = 8;
        } else if (ipelems[0] < 192) {
            // class B
            subnetValue = 16;
        } else if (ipelems[0] < 224) {
            // class C
            subnetValue = 24;
        } else {
            $("#edit-subnet").addClass("is-invalid");
            return;
        }
        $("#edit-subnet").val(subnetValue);
        $("#edit-subnet").addClass("is-valid");
        showNetmask(subnetValue);
    }

    calcBroadcastAddr();
}

function formatIP(e) {
    str = e.target.value.replace(/[^0-9.]/ig, "");
    var ipelems = str.split('.');
    
    newIPstr = ""
    for (i = 0; i < ipelems.length; i++) {
        if (ipelems[i] > 255) ipelems[i] = Math.floor(ipelems[i] / 10);
        newIPstr = newIPstr + String(ipelems[i]) + '.';
    }
    newIPstr = newIPstr.substr(0, newIPstr.length - 1);

    e.target.value = newIPstr.slice(0, 16);
}



function convNetmask(e) {
    if (e.target.value.length == 0) {
        $("#edit-wolnet").val("");

        $("#edit-subnet").removeClass("is-valid");
        $("#edit-subnet").addClass("is-invalid");
        return;
    }

    $("#edit-subnet").removeClass("is-invalid");
    $("#edit-subnet").addClass("is-valid");

    str = e.target.value.replace(/[^0-9]/ig, "");
    if (str > 32) str = "32";
    e.target.value = str.slice(0,2);

    showNetmask(str);
    calcBroadcastAddr();
}

function showNetmask(value) {
    allmasked = Math.floor(Number(value) / 8);
    partmask = Number(value) % 8;

    if ((allmasked == 0) && (partmask == 0)) {
        $("#subnet-display").text("");
        return;
    }
    
    var netmask = "";
    for (i = 0; i < allmasked; i++) {
        netmask += "255."
    }

    if (allmasked < 4) {
        if (partmask == 0) {
            netmask += "0.";
        } else {
            partmask2 = 8 - partmask;
            maskValue = 256 - (1 << partmask2);
            netmask += (String(maskValue) + ".");
        }

        for (i = allmasked; i < 3; i++) {
            netmask += "0.";
        }
    }

    netmask = netmask.substr(0, netmask.length - 1);
    $("#subnet-display").text("Netmask: " + netmask);
}



function calcBroadcastAddr() {
    ipaddr = $("#edit-ip").val();
    netmask = Number($("#edit-subnet").val());

    if ((ipaddr.length == 0) || (netmask.length == 0)) return;

    ipelem = ipaddr.split('.');
    if (netmask > (ipaddr.length * 8)) return;

    var netmaskStr = "";
    for (i = 0; i < ipelem.length; i++) {
        if (netmask == 0) {
            netmaskStr = netmaskStr + "255.";
            continue;
        }
        
        if (netmask > 8) {
            netmaskVal = 8;
            netmask = netmask - 8;
        } else {
            netmaskVal = netmask;
            netmask = 0;
        }
        maskValue = 256 - (1 << (8 - netmaskVal));
        netAria = 255 ^ maskValue;

        netmaskElem = Number(ipelem[i]) & maskValue;
        if (netAria != 0) netmaskElem = netmaskElem | netAria;

        netmaskStr = netmaskStr + String(netmaskElem) + ".";
    }

    netmaskStr = netmaskStr.substr(0, netmaskStr.length - 1);
    $("#edit-wolnet").val(netmaskStr);
}

function checkAvailableFunc() {
    $.ajax({
        type: "GET",
        url: "/api/conf/checkflag?key=readonly&key=allowdownconf"
    }).done(function(availableConf) {
        console.log(availableConf);
        if (availableConf.allowdownconf) {
            $("#download-conf-button").show();
        } else {
            $("#download-conf-button").hide();
        }
        if (availableConf.readonly) {
            $("#machine-add-button").hide();
        } else {
            $("#machine-add-button").show();
        }
    });
}

function downloadConfig() {
    window.location.href = "/api/conf/get"
}

getMachineList();
checkAvailableFunc();

var configID = 0;

var editName = $("#edit-name");
editName.on("keyup", editComputerName)

var editIP = $('#edit-ip');
editIP.on("keyup", formatIPaddr);
var editWolNet = $('#edit-wolnet');
editWolNet.on("keyup", formatIP);

var editMac = $("#edit-mac");
editMac.on("keyup", editMAC);

var editSubnet = $("#edit-subnet");
editSubnet.on("keyup", convNetmask);

