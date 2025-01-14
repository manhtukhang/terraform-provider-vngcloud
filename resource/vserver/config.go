package vserver

var (
	loadBalancerCreating   = []string{"CREATING"}
	loadBalancerCreated    = []string{"CREATED"}
	listenerCreating       = []string{"CREATING"}
	listenerUpdating       = []string{"UPDATING"}
	listenerCreated        = []string{"CREATED"}
	poolCreating           = []string{"CREATING"}
	poolUpdating           = []string{"UPDATING"}
	poolCreated            = []string{"CREATED"}
	networkCreating        = []string{"CREATING"}
	networkCreated         = []string{"ACTIVE"}
	networkDeleting        = []string{"DELETING"}
	networkDeleted         = []string{"DELETED"}
	secgroupRuleCreating   = []string{"CREATING"}
	secgroupRuleCreated    = []string{"ACTIVE"}
	secgroupCreating       = []string{"CREATING"}
	secgroupCreated        = []string{"ACTIVE"}
	serverCreating         = []string{"CREATING", "CREATING-BILLING"}
	serverCreated          = []string{"ACTIVE"}
	serverRebooting        = []string{"REBOOTING"}
	serverRebooted         = []string{"ACTIVE"}
	serverStopping         = []string{"TURNING-OFF"}
	serverStopped          = []string{"STOPPED"}
	serverStarting         = []string{"STARTING"}
	serverStarted          = []string{"ACTIVE"}
	serverResizing         = []string{"CHANGING-FLAVOR", "VERIFYING-FLAVOR"}
	serverResize           = []string{"ACTIVE", "STOPPED"}
	serverDeleting         = []string{"DELETING"}
	serverDeleted          = []string{"DELETED"}
	subnetCreating         = []string{"CREATING"}
	subnetCreated          = []string{"ACTIVE"}
	subnetDeleting         = []string{"DELETING"}
	subnetDeleted          = []string{"DELETED"}
	volumeCreating         = []string{"CREATING", "CREATING-BILLING"}
	volumeCreated          = []string{"AVAILABLE"}
	volumeAttaching        = []string{"AVAILABLE", "ATTACHING"}
	volumeAttached         = []string{"IN-USE"}
	volumeDetaching        = []string{"IN-USE", "DETACHING"}
	volumeResizing         = []string{"RESIZING", "CHANGING-IOPS"}
	volumeResized          = []string{"IN-USE", "AVAILABLE"}
	volumeDetached         = []string{"AVAILABLE", "IN-USE"}
	volumeDeleting         = []string{"DELETING"}
	volumeDeleted          = []string{"DELETED"}
	serverChangingSecGroup = []string{"CHANGING-SECURITY-GROUP"}
	serverChangedSecGroup  = []string{"ACTIVE", "STOPPED"}
)
