package models

import "testing"

type isServiceTest struct {
	serviceName string
	planName    string
	result      bool
}

func TestIsMySQLService(t *testing.T) {
	tests := []isServiceTest{
		{
			"mysql",
			"shared",
			true,
		},
		{
			"somedb",
			"shared-mysql",
			true,
		},
		{
			"aws",
			"rds",
			false,
		},
		{
			"psql",
			"shared",
			false,
		},
	}
	for _, test := range tests {
		serviceInstance := ServiceInstance{
			Service: test.serviceName,
			Plan:    test.planName,
		}
		result := serviceInstance.IsMySQLService()
		if result != test.result {
			t.Errorf("Expected result %v. Real result %v. Data: Service Name '%s' Plan Name '%s'",
				test.result, result, test.serviceName, test.planName)
		}
	}
}

func TestIsPSQLService(t *testing.T) {
	tests := []isServiceTest{
		{
			"psql",
			"shared",
			true,
		},
		{
			"postgres",
			"shared",
			true,
		},
		{
			"somedb",
			"psql",
			true,
		},
		{
			"somedb",
			"postgres",
			true,
		},
		{
			"aws",
			"rds",
			false,
		},
		{
			"mysql",
			"shared",
			false,
		},
	}
	for _, test := range tests {
		serviceInstance := ServiceInstance{
			Service: test.serviceName,
			Plan:    test.planName,
		}
		result := serviceInstance.IsPSQLService()
		if result != test.result {
			t.Errorf("Expected result %v. Real result %v. Data: Service Name '%s' Plan Name '%s'",
				test.result, result, test.serviceName, test.planName)
		}
	}
}
