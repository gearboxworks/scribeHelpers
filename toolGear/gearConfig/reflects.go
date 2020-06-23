package gearConfig


func ReflectGearConfig(ref interface{}) *GearConfig {
	var s *GearConfig

	for range onlyOnce {
		switch ref.(type) {
			case GearConfig:
				s = ref.(*GearConfig)
			//case *GearConfig:
			//	s = (ref.(GearConfig))
			//default:
			//	s = &GearConfig{}
		}
	}

	return s
}


func ReflectGearConfigs(ref interface{}) *GearConfigs {
	var s *GearConfigs

	for range onlyOnce {
		switch ref.(type) {
			case GearConfigs:
				s = ref.(*GearConfigs)
			case map[string]interface{}:
				s = ref.(*GearConfigs)
			//default:
			//	s = &GearConfigs{}
		}
	}

	return s
}
