package componenttest

import (
	"math"
	"reflect"
	"testing"
	"time"
	"xenapi"
)

func TestVMGetAllRecords(t *testing.T) {
	session, err := GetSession("xapi-24/vm_get_all_records_02")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	var vmRecord0 = xenapi.VMRecord{
		UUID: "c5db1298-3e97-53bb-4f97-bd039ca8c217",
		AllowedOperations: []xenapi.VMOperations{
			"changing_dynamic_range",
			"migrate_send",
			"pool_migrate",
			"changing_VCPUs_live",
			"suspend",
			"hard_reboot",
			"hard_shutdown",
			"clean_reboot",
			"clean_shutdown",
			"pause",
			"checkpoint",
			"snapshot",
		},
		CurrentOperations:       map[string]xenapi.VMOperations{},
		NameLabel:               "win11-x64",
		NameDescription:         "win11-x64",
		PowerState:              "Running",
		UserVersion:             1,
		IsATemplate:             false,
		IsDefaultTemplate:       false,
		SuspendVDI:              "OpaqueRef:NULL",
		ResidentOn:              "OpaqueRef:4a0fde9c-5709-4f66-9968-64fb6162bd97",
		ScheduledToBeResidentOn: "OpaqueRef:NULL",
		Affinity:                "OpaqueRef:NULL",
		MemoryOverhead:          71303168,
		MemoryTarget:            8589934592,
		MemoryStaticMax:         8589934592,
		MemoryDynamicMax:        8589934592,
		MemoryDynamicMin:        8589934592,
		MemoryStaticMin:         4294967296,
		VCPUsParams: map[string]string{
			"weight": "256",
		},
		VCPUsMax:               2,
		VCPUsAtStartup:         2,
		ActionsAfterSoftreboot: "soft_reboot",
		ActionsAfterShutdown:   "destroy",
		ActionsAfterReboot:     "restart",
		ActionsAfterCrash:      "preserve",
		Consoles: []xenapi.ConsoleRef{
			"OpaqueRef:87122f55-ce67-7a33-5d93-0d25206e83a0",
		},
		VIFs: []xenapi.VIFRef{
			"OpaqueRef:5f1d3810-45d1-0762-81f6-18b198b52e0c",
		},
		VBDs: []xenapi.VBDRef{
			"OpaqueRef:f7ef1069-2538-089a-2a77-e42e3b9e336c",
			"OpaqueRef:c92a38a8-87aa-2d04-622f-84a788d9f9ef",
			"OpaqueRef:e3e47cc5-5752-df57-b3a3-89bcc7336397",
		},
		VUSBs:      []xenapi.VUSBRef{},
		CrashDumps: []xenapi.CrashdumpRef{},
		VTPMs: []xenapi.VTPMRef{
			"OpaqueRef:60f5a19e-773b-f831-c9c2-35f07b632603",
		},
		PVBootloader:     "",
		PVKernel:         "",
		PVRamdisk:        "",
		PVArgs:           "",
		PVBootloaderArgs: "",
		PVLegacyArgs:     "",
		HVMBootPolicy:    "BIOS order",
		HVMBootParams: map[string]string{
			"firmware": "uefi",
			"order":    "cd",
		},
		HVMShadowMultiplier: 1.0,
		Platform: map[string]string{
			"acpi":                    "1",
			"acpi_laptop_slate":       "1",
			"apic":                    "true",
			"cores-per-socket":        "2",
			"device-model":            "qemu-upstream-uefi",
			"device_id":               "0002",
			"hpet":                    "true",
			"nx":                      "true",
			"pae":                     "true",
			"secureboot":              "true",
			"timeoffset":              "0",
			"vga":                     "std",
			"videoram":                "8",
			"viridian":                "true",
			"viridian_apic_assist":    "true",
			"viridian_crash_ctl":      "true",
			"viridian_reference_tsc":  "true",
			"viridian_stimer":         "true",
			"viridian_time_ref_count": "true",
			"vtpm":                    "true",
		},
		PCIBus: "",
		OtherConfig: map[string]string{
			"base_template_name": "Windows 11",
			"import_task":        "OpaqueRef:06641b68-0695-5bfa-2fce-ec2c1c914097",
			"install-methods":    "cdrom",
			"mac_seed":           "8414a948-97f5-2a8b-2070-c26b95e5d578",
			"xenrt-distro":       "win11-x64",
		},
		Domid:   50,
		Domarch: "",
		LastBootCPUFlags: map[string]string{
			"features": "1fcbfbff-f7fa3203-2c100800-00000021-00000001-000007ab-00000000-00000000-00101000-bc000400-00000000-00000000-00000000-00000000-00000000-00000000-04000000-00000000-00000000-00000000-00000000-00000000",
			"vendor":   "GenuineIntel",
		},
		IsControlDomain:  false,
		Metrics:          "OpaqueRef:29a46db3-9b13-7df1-788c-638c194d8cf1",
		GuestMetrics:     "OpaqueRef:60d05d5e-4888-ebeb-f070-c57f90410417",
		LastBootedRecord: "",
		Recommendations:  "<restrictions><restriction field=\"memory-static-max\" max=\"1649267441664\"/><restriction field=\"vcpus-min\" min=\"2\"/><restriction field=\"vcpus-max\" max=\"64\"/><restriction field=\"has-vendor-device\" value=\"true\"/><restriction field=\"supports-bios\" value=\"no\"/><restriction field=\"supports-uefi\" value=\"yes\"/><restriction field=\"supports-secure-boot\" value=\"yes\"/><restriction max=\"255\" property=\"number-of-vbds\"/><restriction max=\"7\" property=\"number-of-vifs\"/></restrictions>",
		XenstoreData: map[string]string{
			"vm-data":                "",
			"vm-data/mmio-hole-size": "268435456",
		},
		HaAlwaysRun:             false,
		HaRestartPriority:       "",
		IsASnapshot:             false,
		SnapshotOf:              "OpaqueRef:NULL",
		Snapshots:               []xenapi.VMRef{},
		SnapshotTime:            time.Unix(1627478400, 0),
		TransportableSnapshotID: "",
		Blobs:                   map[string]xenapi.BlobRef{},
		Tags:                    []string{},
		BlockedOperations:       map[xenapi.VMOperations]string{},
		SnapshotInfo:            map[string]string{},
		SnapshotMetadata:        "",
		Parent:                  "OpaqueRef:NULL",
		Children:                []xenapi.VMRef{},
		BiosStrings: map[string]string{
			"baseboard-asset-tag":           "",
			"baseboard-location-in-chassis": "",
			"baseboard-manufacturer":        "",
			"baseboard-product-name":        "",
			"baseboard-serial-number":       "",
			"baseboard-version":             "",
			"bios-vendor":                   "Xen",
			"bios-version":                  "",
			"enclosure-asset-tag":           "",
			"hp-rombios":                    "",
			"oem-1":                         "Xen",
			"oem-2":                         "MS_VM_CERT/SHA1/bdbeb6e0a816d43fa6d3fe8aaef04c2bad9d3e3d",
			"system-manufacturer":           "Xen",
			"system-product-name":           "HVM domU",
			"system-serial-number":          "",
			"system-version":                "",
		},
		ProtectionPolicy:            "OpaqueRef:NULL",
		IsSnapshotFromVmpp:          false,
		SnapshotSchedule:            "OpaqueRef:NULL",
		IsVmssSnapshot:              false,
		Appliance:                   "OpaqueRef:NULL",
		StartDelay:                  0,
		ShutdownDelay:               0,
		Order:                       0,
		VGPUs:                       []xenapi.VGPURef{},
		AttachedPCIs:                []xenapi.PCIRef{},
		SuspendSR:                   "OpaqueRef:5ec2d621-7b62-f571-d38b-754341a973e5",
		Version:                     0,
		GenerationID:                "209202393574174067:5365773919542227445",
		HardwarePlatformVersion:     2,
		HasVendorDevice:             false,
		RequiresReboot:              false,
		ReferenceLabel:              "windows-11",
		DomainType:                  "hvm",
		NVRAM:                       map[string]string{"EFI-variables": "VkFSUwIAAAAsAAAAA"},
		PendingGuidances:            []xenapi.UpdateGuidances{},
		PendingGuidancesRecommended: []xenapi.UpdateGuidances{},
		PendingGuidancesFull:        []xenapi.UpdateGuidances{},
	}
	var vmRecord1 = xenapi.VMRecord{
		UUID: "060aee48-7ded-d2bd-6f0f-f224bd99c960",
		AllowedOperations: []xenapi.VMOperations{
			"changing_dynamic_range",
			"migrate_send",
			"pool_migrate",
			"changing_VCPUs_live",
			"suspend",
			"hard_reboot",
			"hard_shutdown",
			"clean_reboot",
			"clean_shutdown",
			"pause",
			"checkpoint",
			"snapshot",
		},
		CurrentOperations: map[string]xenapi.VMOperations{},
		NameLabel:         "ubuntu2204",
		IsATemplate:       false,
		SuspendVDI:        "OpaqueRef:NULL",
		ResidentOn:        "OpaqueRef:4a0fde9c-5709-4f66-9968-64fb6162bd97",
		MemoryStaticMax:   4294967296,
		VBDs: []xenapi.VBDRef{
			"OpaqueRef:4a54632c-a6b1-5535-4ef4-0d36d85e07df",
			"OpaqueRef:56d316f7-e8d0-54f0-2492-a0952c6946c4",
		},
		HVMShadowMultiplier: 1.0,
		Platform: map[string]string{
			"acpi":         "1",
			"apic":         "true",
			"device-model": "qemu-upstream-compat",
			"device_id":    "0001",
			"hpet":         "true",
			"nx":           "true",
			"pae":          "true",
			"timeoffset":   "0",
			"vga":          "std",
			"videoram":     "8",
			"viridian":     "false",
		},
		Domid:        45,
		SnapshotTime: time.Date(2023, time.July, 29, 13, 20, 01, 0, time.UTC),
	}
	var expectedResult = map[xenapi.VMRef]xenapi.VMRecord{
		"OpaqueRef:6ef08bce-0bf0-30ff-804f-5f0ee4bbdd13": vmRecord0,
		"OpaqueRef:7cdfad16-7e0e-411f-1b6c-4308fd33b164": vmRecord1,
	}
	result, err := xenapi.VM.GetAllRecords(session)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Log("The result returned not the same with expected -> The XAPI outcome diverges from the anticipated value.")
		t.Fail()
		return
	}

}

func TestVMImport(t *testing.T) {
	session, err := GetSession("xapi-24/vm_import_03")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	var srRef xenapi.SRRef = "OpaqueRef:5ec2d621-7b62-f571-s38r-754341a973e5"
	var VMImportURL = "http://vm_import_url"

	var expectedResult = []xenapi.VMRef{"OpaqueRef:5ec2d621-7b62-f571-v38m-754341a973e5"}
	result, err := xenapi.VM.Import(session, VMImportURL, srRef, false, false)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Log("The result returned not the same with expected -> The XAPI outcome diverges from the anticipated value.")
		t.Fail()
		return
	}

}

func TestVMRetrieveWlbRecommendations(t *testing.T) {
	session, err := GetSession("xapi-24/vm_retrieve_wlb_recommendations_04")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	var vmRef xenapi.VMRef = "OpaqueRef:5ec2d621-7b62-f571-v38m-754341a973e5"
	expectedResult := map[xenapi.HostRef][]string{
		"OpaqueRef:4a0fde9c-5709-4f66-9968-64fb6162bd97": []string{
			"host0",
		},
	}
	result, err := xenapi.VM.RetrieveWlbRecommendations(session, vmRef)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Log("The result returned not the same with expected -> The XAPI outcome diverges from the anticipated value.")
		t.Fail()
		return
	}
}

func TestVMMigrateSend(t *testing.T) {
	session, err := GetSession("xapi-24/vm_migrate_send_05")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	var vmRef xenapi.VMRef = "OpaqueRef:5vm2d621-7b62-f571-vm00-754341a973e5"
	dest := map[string]string{
		"name_label": "host1",
	}
	vdiMap := map[xenapi.VDIRef]xenapi.SRRef{
		"OpaqueRef:5vdid621-7b62-f571-vdi8-754341a973e5": "OpaqueRef:5sr2d621-7b62-f571-s38r-754341a973e5",
	}
	vifMap := map[xenapi.VIFRef]xenapi.NetworkRef{
		"OpaqueRef:9vifb8a9-f6a3-69d2-vifb-7a7b7c81c49e": "OpaqueRef:9network-f6a3-69d2-netw-7a7b7c81c49e",
	}
	var options map[string]string
	vgpuMap := map[xenapi.VGPURef]xenapi.GPUGroupRef{
		"OpaqueRef:9vgpu8a9-f6a3-69d2-vgpu-7a7b7c81c49e": "OpaqueRef:9gpugroup-f6a3-69d2-grou-7a7b7c81c49e",
	}

	var expectedResult xenapi.VMRef = "OpaqueRef:5vm2d621-7b62-f571-vm01-754341a973e5"
	result, err := xenapi.VM.MigrateSend(session, vmRef, dest, true, vdiMap, vifMap, options, vgpuMap)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if result != expectedResult {
		t.Log("The result returned not the same with expected -> The XAPI outcome diverges from the anticipated value.")
		t.Fail()
		return
	}
}

func TestVMGetDataSources(t *testing.T) {
	session, err := GetSession("xapi-24/vm_get_data_sources_06")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	var vmRef xenapi.VMRef = "OpaqueRef:6ef08bce-0bf0-30ff-804f-5f0ee4bbdd13"
	var expectedResult = []xenapi.DataSourceRecord{
		{
			NameLabel:       "memory",
			NameDescription: "Memory currently allocated to VM",
			Enabled:         true,
			Standard:        true,
			Units:           "B",
			Min:             0.0,
			Max:             math.Inf(1),
			Value:           0.0,
		},
	}
	result, err := xenapi.VM.GetDataSources(session, vmRef)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Log("The result returned not the same with expected -> The XAPI outcome diverges from the anticipated value.")
		t.Fail()
		return
	}
}

func TestVMGetSnapshotTime(t *testing.T) {
	session, err := GetSession("xapi-24/vm_get_snapshot_time_07")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	var vmRef xenapi.VMRef = "OpaqueRef:6ef08bce-0bf0-30ff-804f-5f0ee4bbdd13"
	var expectedResult time.Time = time.Date(2024, time.April, 20, 12, 0, 0, 0, time.UTC)
	result, err := xenapi.VM.GetSnapshotTime(session, vmRef)

	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if result != expectedResult {
		t.Log("The result returned not the same with expected -> The XAPI outcome diverges from the anticipated value.")
		t.Fail()
		return
	}
}

func TestVMQueryDataSource(t *testing.T) {
	session, err := GetSession("xapi-24/vm_query_data_source_08")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	var vmRef xenapi.VMRef = "OpaqueRef:6ef08bce-0bf0-30ff-804f-5f0ee4bbdd13"

	result, err := xenapi.VM.QueryDataSource(session, vmRef, "CPU0 usage")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	if !math.IsNaN(result) {
		t.Log("The result returned not the same with expected -> The XAPI outcome diverges from the anticipated value.")
		t.Fail()
		return
	}
}

func TestVMAsyncImport(t *testing.T) {
	session, err := GetSession("xapi-24/async_vm_import_12")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	var srRef xenapi.SRRef = "OpaqueRef:5ec2d621-7b62-f571-d38b-754341a973e5"
	var VMImportURL = "http://vm_import_url"
	var expectedResult xenapi.TaskRef = "OpaqueRef:5task621-7b62-f571-d38b-754341a973e5"
	result, err := xenapi.VM.AsyncImport(session, VMImportURL, srRef, false, false)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if result != expectedResult {
		t.Log("The expected error is not the same with the returned one!")
		t.Fail()
		return
	}
}
