// Peripheral: ADC_Common_Periph  Analog to Digital Converter.
// Instances:
//  ADC1_2  mmap.ADC1_2_BASE
//  ADC3_4  mmap.ADC3_4_BASE
// Registers:
//  0x00 32  CSR ADC Common status register.
//  0x08 32  CCR ADC common control register.
//  0x0C 32  CDR ADC common regular data register for dual.
// Import:
//  stm32/o/f303xe/mmap
package adc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	ADRDY_MST       CSR_Bits = 0x01 << 0  //+ Master ADC ready.
	ADRDY_EOSMP_MST CSR_Bits = 0x01 << 1  //+ End of sampling phase flag of the master ADC.
	ADRDY_EOC_MST   CSR_Bits = 0x01 << 2  //+ End of regular conversion of the master ADC.
	ADRDY_EOS_MST   CSR_Bits = 0x01 << 3  //+ End of regular sequence flag of the master ADC.
	ADRDY_OVR_MST   CSR_Bits = 0x01 << 4  //+ Overrun flag of the master ADC.
	ADRDY_JEOC_MST  CSR_Bits = 0x01 << 5  //+ End of injected conversion of the master ADC.
	ADRDY_JEOS_MST  CSR_Bits = 0x01 << 6  //+ End of injected sequence flag of the master ADC.
	AWD1_MST        CSR_Bits = 0x01 << 7  //+ Analog watchdog 1 flag of the master ADC.
	AWD2_MST        CSR_Bits = 0x01 << 8  //+ Analog watchdog 2 flag of the master ADC.
	AWD3_MST        CSR_Bits = 0x01 << 9  //+ Analog watchdog 3 flag of the master ADC.
	JQOVF_MST       CSR_Bits = 0x01 << 10 //+ Injected context queue overflow flag of the master ADC.
	ADRDY_SLV       CSR_Bits = 0x01 << 16 //+ Slave ADC ready.
	ADRDY_EOSMP_SLV CSR_Bits = 0x01 << 17 //+ End of sampling phase flag of the slave ADC.
	ADRDY_EOC_SLV   CSR_Bits = 0x01 << 18 //+ End of regular conversion of the slave ADC.
	ADRDY_EOS_SLV   CSR_Bits = 0x01 << 19 //+ End of regular sequence flag of the slave ADC.
	ADRDY_OVR_SLV   CSR_Bits = 0x01 << 20 //+ Overrun flag of the slave ADC.
	ADRDY_JEOC_SLV  CSR_Bits = 0x01 << 21 //+ End of injected conversion of the slave ADC.
	ADRDY_JEOS_SLV  CSR_Bits = 0x01 << 22 //+ End of injected sequence flag of the slave ADC.
	AWD1_SLV        CSR_Bits = 0x01 << 23 //+ Analog watchdog 1 flag of the slave ADC.
	AWD2_SLV        CSR_Bits = 0x01 << 24 //+ Analog watchdog 2 flag of the slave ADC.
	AWD3_SLV        CSR_Bits = 0x01 << 25 //+ Analog watchdog 3 flag of the slave ADC.
	JQOVF_SLV       CSR_Bits = 0x01 << 26 //+ Injected context queue overflow flag of the slave ADC.
)

const (
	ADRDY_MSTn       = 0
	ADRDY_EOSMP_MSTn = 1
	ADRDY_EOC_MSTn   = 2
	ADRDY_EOS_MSTn   = 3
	ADRDY_OVR_MSTn   = 4
	ADRDY_JEOC_MSTn  = 5
	ADRDY_JEOS_MSTn  = 6
	AWD1_MSTn        = 7
	AWD2_MSTn        = 8
	AWD3_MSTn        = 9
	JQOVF_MSTn       = 10
	ADRDY_SLVn       = 16
	ADRDY_EOSMP_SLVn = 17
	ADRDY_EOC_SLVn   = 18
	ADRDY_EOS_SLVn   = 19
	ADRDY_OVR_SLVn   = 20
	ADRDY_JEOC_SLVn  = 21
	ADRDY_JEOS_SLVn  = 22
	AWD1_SLVn        = 23
	AWD2_SLVn        = 24
	AWD3_SLVn        = 25
	JQOVF_SLVn       = 26
)

const (
	MULTI    CCR_Bits = 0x1F << 0  //+ Multi ADC mode selection.
	MULTI_0  CCR_Bits = 0x01 << 0  //  MULTI bit 0.
	MULTI_1  CCR_Bits = 0x02 << 0  //  MULTI bit 1.
	MULTI_2  CCR_Bits = 0x04 << 0  //  MULTI bit 2.
	MULTI_3  CCR_Bits = 0x08 << 0  //  MULTI bit 3.
	MULTI_4  CCR_Bits = 0x10 << 0  //  MULTI bit 4.
	DELAY    CCR_Bits = 0x0F << 8  //+ Delay between 2 sampling phases.
	DELAY_0  CCR_Bits = 0x01 << 8  //  DELAY bit 0.
	DELAY_1  CCR_Bits = 0x02 << 8  //  DELAY bit 1.
	DELAY_2  CCR_Bits = 0x04 << 8  //  DELAY bit 2.
	DELAY_3  CCR_Bits = 0x08 << 8  //  DELAY bit 3.
	MDMACFG  CCR_Bits = 0x01 << 13 //+ DMA configuration for multi-ADC mode.
	MDMA     CCR_Bits = 0x03 << 14 //+ DMA mode for multi-ADC mode.
	MDMA_0   CCR_Bits = 0x01 << 14 //  MDMA bit 0.
	MDMA_1   CCR_Bits = 0x02 << 14 //  MDMA bit 1.
	CKMODE   CCR_Bits = 0x03 << 16 //+ ADC clock mode.
	CKMODE_0 CCR_Bits = 0x01 << 16 //  CKMODE bit 0.
	CKMODE_1 CCR_Bits = 0x02 << 16 //  CKMODE bit 1.
	VREFEN   CCR_Bits = 0x01 << 22 //+ VREFINT enable.
	TSEN     CCR_Bits = 0x01 << 23 //+ Temperature sensor enable.
	VBATEN   CCR_Bits = 0x01 << 24 //+ VBAT enable.
)

const (
	MULTIn   = 0
	DELAYn   = 8
	MDMACFGn = 13
	MDMAn    = 14
	CKMODEn  = 16
	VREFENn  = 22
	TSENn    = 23
	VBATENn  = 24
)

const (
	RDATA_MST    CDR_Bits = 0xFFFF << 0  //+ Regular Data of the master ADC.
	RDATA_MST_0  CDR_Bits = 0x01 << 0    //  RDATA_MST bit 0.
	RDATA_MST_1  CDR_Bits = 0x02 << 0    //  RDATA_MST bit 1.
	RDATA_MST_2  CDR_Bits = 0x04 << 0    //  RDATA_MST bit 2.
	RDATA_MST_3  CDR_Bits = 0x08 << 0    //  RDATA_MST bit 3.
	RDATA_MST_4  CDR_Bits = 0x10 << 0    //  RDATA_MST bit 4.
	RDATA_MST_5  CDR_Bits = 0x20 << 0    //  RDATA_MST bit 5.
	RDATA_MST_6  CDR_Bits = 0x40 << 0    //  RDATA_MST bit 6.
	RDATA_MST_7  CDR_Bits = 0x80 << 0    //  RDATA_MST bit 7.
	RDATA_MST_8  CDR_Bits = 0x100 << 0   //  RDATA_MST bit 8.
	RDATA_MST_9  CDR_Bits = 0x200 << 0   //  RDATA_MST bit 9.
	RDATA_MST_10 CDR_Bits = 0x400 << 0   //  RDATA_MST bit 10.
	RDATA_MST_11 CDR_Bits = 0x800 << 0   //  RDATA_MST bit 11.
	RDATA_MST_12 CDR_Bits = 0x1000 << 0  //  RDATA_MST bit 12.
	RDATA_MST_13 CDR_Bits = 0x2000 << 0  //  RDATA_MST bit 13.
	RDATA_MST_14 CDR_Bits = 0x4000 << 0  //  RDATA_MST bit 14.
	RDATA_MST_15 CDR_Bits = 0x8000 << 0  //  RDATA_MST bit 15.
	RDATA_SLV    CDR_Bits = 0xFFFF << 16 //+ Regular Data of the master ADC.
	RDATA_SLV_0  CDR_Bits = 0x01 << 16   //  RDATA_SLV bit 0.
	RDATA_SLV_1  CDR_Bits = 0x02 << 16   //  RDATA_SLV bit 1.
	RDATA_SLV_2  CDR_Bits = 0x04 << 16   //  RDATA_SLV bit 2.
	RDATA_SLV_3  CDR_Bits = 0x08 << 16   //  RDATA_SLV bit 3.
	RDATA_SLV_4  CDR_Bits = 0x10 << 16   //  RDATA_SLV bit 4.
	RDATA_SLV_5  CDR_Bits = 0x20 << 16   //  RDATA_SLV bit 5.
	RDATA_SLV_6  CDR_Bits = 0x40 << 16   //  RDATA_SLV bit 6.
	RDATA_SLV_7  CDR_Bits = 0x80 << 16   //  RDATA_SLV bit 7.
	RDATA_SLV_8  CDR_Bits = 0x100 << 16  //  RDATA_SLV bit 8.
	RDATA_SLV_9  CDR_Bits = 0x200 << 16  //  RDATA_SLV bit 9.
	RDATA_SLV_10 CDR_Bits = 0x400 << 16  //  RDATA_SLV bit 10.
	RDATA_SLV_11 CDR_Bits = 0x800 << 16  //  RDATA_SLV bit 11.
	RDATA_SLV_12 CDR_Bits = 0x1000 << 16 //  RDATA_SLV bit 12.
	RDATA_SLV_13 CDR_Bits = 0x2000 << 16 //  RDATA_SLV bit 13.
	RDATA_SLV_14 CDR_Bits = 0x4000 << 16 //  RDATA_SLV bit 14.
	RDATA_SLV_15 CDR_Bits = 0x8000 << 16 //  RDATA_SLV bit 15.
)

const (
	RDATA_MSTn = 0
	RDATA_SLVn = 16
)