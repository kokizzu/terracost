package usage

const (
	// Key is the key used to set the usage
	// on the values passed to the resources
	Key string = "tc_usage"
)

// Default is the default Usage that will be used if none is configured
var Default = Usage{
	ResourceDefaultTypeUsage: map[string]interface{}{
		// AWS
		"aws_cloudwatch_log_group": map[string]interface{}{
			"storage_gb":                       200,
			"monthly_data_ingested_gb":         10,
			"monthly_data_scanned_insights_gb": 20,
		},
		"aws_eks_node_group": map[string]interface{}{
			"instances":                        15,
			"operating_system":                 "linux",
			"reserved_instance_type":           "standard",
			"reserved_instance_term":           "1_year",
			"reserved_instance_payment_option": "partial_upfront",
			"monthly_cpu_credit_hrs":           350,
			"vcpu_count":                       2,
		},
		"aws_efs_file_system": map[string]interface{}{
			"storage_gb":                         180,
			"infrequent_access_storage_gb":       10,
			"monthly_infrequent_access_read_gb":  20,
			"monthly_infrequent_access_write_gb": 30,
		},
		"aws_fsx_openzfs_file_system": map[string]interface{}{
			"backup_storage_gb": 1024,
		},
		"aws_fsx_windows_file_system": map[string]interface{}{
			"backup_storage_gb": 1024,
		},
		"aws_fsx_ontap_file_system": map[string]interface{}{
			"backup_storage_gb": 1024,
		},
		"aws_fsx_lustre_file_system": map[string]interface{}{
			"backup_storage_gb": 1024,
		},
		"aws_nat_gateway": map[string]interface{}{
			"monthly_data_processed_gb": 10,
		},
		"aws_rds_cluster": map[string]interface{}{
			"capacity_units_per_hr":        0.5,
			"storage_gb":                   50,
			"write_requests_per_sec":       4,
			"read_requests_per_sec":        4,
			"backup_snapshot_size_gb":      60,
			"average_statements_per_hr":    2500,
			"change_records_per_statement": 0.095,
			"backtrack_window_hrs":         150,
			"snapshot_export_size_gb":      300,
		},
		"aws_rds_cluster_instance": map[string]interface{}{
			"monthly_additional_performance_insights_requests": 500000,
			"capacity_units_per_hr":                            0.5,
		},
		"aws_s3_bucket": map[string]interface{}{
			"storage_gb":               200,
			"monthly_outbound_data_gb": 10,
		},
		"aws_s3_bucket_analytics_configuration": map[string]interface{}{
			"monthly_monitored_objects": 50000000,
		},

		"aws_s3_bucket_inventory": map[string]interface{}{
			"monthly_listed_objects": 2000000000,
		},
		"aws_secretsmanager_secret": map[string]interface{}{
			"monthly_requests": 1000000,
		},
		"aws_sqs_queue": map[string]interface{}{
			"monthly_requests": 15000000,
			"request_size_kb":  16,
		},

		// Azure
		"azurerm_bastion_host": map[string]interface{}{
			"monthly_outbound_data_gb": 40,
		},
		"azurerm_nat_gateway": map[string]interface{}{
			"monthly_data_processed_gb": 150,
		},
		"azurerm_virtual_network_gateway": map[string]interface{}{
			"monthly_data_transfer_gb": 150,
		},
		"azurerm_managed_disk": map[string]interface{}{
			// Number of disk operations (writes, reads, deletes)
			"monthly_disk_operations": 100000000,
		},
		"azurerm_virtual_machine": map[string]interface{}{
			"os_disk": map[string]interface{}{
				// Number of disk operations (writes, reads, deletes)
				"monthly_disk_operations": 100000000,
			},
		},
		"azurerm_linux_virtual_machine": map[string]interface{}{
			"os_disk": map[string]interface{}{
				// Number of disk operations (writes, reads, deletes)
				"monthly_disk_operations": 100000000,
			},
		},
		"azurerm_windows_virtual_machine": map[string]interface{}{
			"os_disk": map[string]interface{}{
				// Number of disk operations (writes, reads, deletes)
				"monthly_disk_operations": 100000000,
			},
		},
		"azurerm_storage_share": map[string]interface{}{
			"monthly_write_transactions": 1000000,
			"monthly_list_transactions":  1000000,
			"monthly_read_transactions":  1000000,
			"monthly_other_transactions": 1000000,
		},
		"azurerm_public_ip": map[string]interface{}{
			"monthly_hours": 730, // Corresponds to a full month
		},
		"azurerm_private_endpoint": map[string]interface{}{
			"monthly_hours": 730, // Corresponds to a full month
		},
		"azurerm_postgresql_flexible_server": map[string]interface{}{
			"additional_backup_storage_gb": 50,  // in GB
			"monthly_hours":                730, // Corresponds to a full month
		},
	},
}

// Usage is the struct defining all the configure usages
type Usage struct {
	ResourceDefaultTypeUsage map[string]interface{} `json:"resource_default_type_usage" yaml:"resource_default_type_usage"`
}

// GetUsage will return the usage from the resource rt (ex: aws_instance)
func (u Usage) GetUsage(rt string) map[string]interface{} {
	us, ok := u.ResourceDefaultTypeUsage[rt]
	if ok {
		return us.(map[string]interface{})
	}

	return nil
}
