package logitech

type DIJOYSTATE2 struct {
	LX         int32
	LY         int32
	LZ         int32
	LRx        int32
	LRy        int32
	LRz        int32
	RglSlider  [2]int32
	RgdwPOV    [4]int32
	RgbButtons [128]byte
	LVX        int32
	LVY        int32
	LVZ        int32
	LVRx       int32
	LVRy       int32
	LVRz       int32
	RglVSlider [2]int32
	LAX        int32
	LAY        int32
	LAZ        int32
	LARx       int32
	LARy       int32
	LARz       int32
	RglASlider [2]int32
	LFX        int32
	LFY        int32
	LFZ        int32
	LFRx       int32
	LFRy       int32
	LFRz       int32
	RglFSlider [2]int32
}
type DIJOYSTATE2ENGINES struct {
	lX         int32     /* x-axis position              */
	lY         int32     /* y-axis position              */
	lZ         int32     /* z-axis position              */
	lRx        int32     /* x-axis rotation              */
	lRy        int32     /* y-axis rotation              */
	lRz        int32     /* z-axis rotation              */
	rglSlider  [2]int32  /* extra axes positions         */
	rgdwPOV    [4]int32  /* POV directions               */
	rgbButtons [128]byte /* 128 buttons                  */
	lVX        int32     /* x-axis velocity              */
	lVY        int32     /* y-axis velocity              */
	lVZ        int32     /* z-axis velocity              */
	lVRx       int32     /* x-axis angular velocity      */
	lVRy       int32     /* y-axis angular velocity      */
	lVRz       int32     /* z-axis angular velocity      */
	rglVSlider [2]int32  /* extra axes velocities        */
	lAX        int32     /* x-axis acceleration          */
	lAY        int32     /* y-axis acceleration          */
	lAZ        int32     /* z-axis acceleration          */
	lARx       int32     /* x-axis angular acceleration  */
	lARy       int32     /* y-axis angular acceleration  */
	lARz       int32     /* z-axis angular acceleration  */
	rglASlider [2]int32  /* extra axes accelerations     */
	lFX        int32     /* x-axis force                 */
	lFY        int32     /* y-axis force                 */
	lFZ        int32     /* z-axis force                 */
	lFRx       int32     /* x-axis torque                */
	lFRy       int32     /* y-axis torque                */
	lFRz       int32     /* z-axis torque                */
	rglFSlider [2]int32  /* extra axes forces            */
}

type LogiControllerPropertiesData struct {
	forceEnable          bool
	overallGain          int32
	springGain           int32
	damperGain           int32
	defaultSpringEnabled bool
	defaultSpringGain    int32
	combinePedals        bool
	wheelRange           int32
	gameSettingsEnabled  bool
	allowGameSettings    bool
}
