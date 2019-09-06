package charge

import svg "github.com/ajstarks/svgo"

var EaglesHeadErased = Charge{
	Identifier: "eagles-head-erased",
	Name:       "eagle's head erased",
	Noun:       "eagle's head",
	NounPlural: "eagle's heads",
	Descriptor: "erased",
	SingleOnly: false,
	Tags: []string{
		"animal",
		"bird",
		"eagle",
		"pride",
		"strong",
	},
	Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
		pathData0 := "M252.1,118.1c-0.1,3.6,1.4,6.3,2.1,9.2c2.4,10.4-2.4,19.7-3.6,29.6c-0.4,3-1.2,6-1.5,9.1c-0.1,2-1.3,3.3-2.7,3.8    c-1.8,0.7-2.2-1.5-2.8-2.7c-1-1.9-2-3.9-3.1-5.9c-4.6,8-5.9,17-8.4,25.6c-1,3.3,0.6,6.9,2.8,9.5c2,2.3,3.3,4.6,1.2,6.9    c-2.1,2.2-3.8-1.1-5.9-1.6c-0.8-0.2-1.3-1.8-2.6-2.5c-0.2,6.3-2.8,11.3-5.7,16.3c-2.6,4.5-0.3,14.3,4.2,19.4    c5.1,5.8,5.6,5.8,10.1-0.3c2.4,4.3,2.4,6.1-0.5,10.2c-1.2,1.6-0.4,1.8,0.8,2.8c3.6,3,4,8,6.3,11.8c0.4,0.6,0.1,2.1-0.4,2.8    c-3.4,4.2-3.9,9.4-5.5,14.2c-1.5,4.5-9.8,9.1-14.1,7.7c-2.3-0.7-4.8-2.2-4.8-5c-0.1-3,3.1-1.8,4.7-2.8c0.3-0.2,1,0,1.7,0    c0.6,1.6-2.3,2.1-1.1,3.6c0.2,0.3,1.1,0.4,1.3,0.2c1-1,2.2-1.7,2.5-3.5c0.9-6.8-2.1-10.7-8.8-12.4c-12.7-3.2-25.7-5.2-36.7-13.5    c-5.1-3.8-7.6-8.5-9.5-14.1c-0.3-0.9-0.7-1.4-1.5-2.1c-5.5,10-13.1,18.6-18.1,28.9c-3.4,7.1-1.8,17.2,3.8,22.7    c4.2,4.1,7.7,3.2,9.8-2.1c1.2-2.9,3.1-1.9,4.8-0.2c4.8,4.8,2.1,12.7-5.6,14.4c-13.4,3.1-26,1.7-35.6-9.7    c-3.6-4.3-4.5-9.8-6.8-14.8c-2.7-6.1-2.5-12.6-3.9-18.8c-1.5-6.7-3.3-7.5-9.7-4.6c-5.8,2.6-11.2,6.1-17.4,8.1    c-2.6,0.8-3.9,3.3-5.5,5.2c-1.2,1.4-0.5,3.4,1,4.8c1.9,1.8,3.2,0.9,4.6-0.6c1.1-1.2,2.3-2.2,3.9-3.8c-0.3,5.7-3.9,8.5-7.9,9    c-3.3,0.4-7.3-1.4-9.1-5.1c-4-8.5-11.6-10.9-19.9-12.6c-3.6-0.7-7.1-2-10.7-2.8c-1.2-0.3-2.4-2.1-3.9-0.2    c-1.7,2.2,0.7,2.9,1.5,4.2c-6,1.9-8.9,1-11.8-3.4c-1.8-2.8-1.9-6,0.7-7.9c2.6-1.9,2-3.3,1.1-5.6c-3.3-8.7-4-18-1.9-26.8    c2.1-8.7,8.7-14.2,17.1-17.3c2.9-1,5.8-1.9,8.6-3c2.8-1.1,5.8-2,5.3-6.2c-0.1-0.6,0.4-3.1,2.1-2.6c1.6,0.6,3.4,2,3.6,3.8    c0.5,4.2,0.1,8-4.5,10.3c-4.3,2.2-8.2,5.1-12.1,8c-2.8,2-3.2,7.9-1,11.3c2.9,4.4,6.4,8.1,11.4,10.3c4.1,1.8,7.2,1,9.3-3    c1.8-3.4,3.4-6.9,5.1-10.4c-1.7-0.9-4.7-0.7-4.7-2.5c0-1.9,2.1-3.4,4.3-4.2c4.2-1.7,6.3-5.6,8.2-9.3c2.1-4,3.8-8.1,6.1-12    c1.2-2.1,5.8-4.4,1.6-8c-0.3-0.2,2.1-3.3,3.2-5.1c-2.2-1.9-7.2-1.1-7.2-4.3c-0.1-2.2,4.5-2.7,6.9-4.2c1.8-1.1,3.6-2.3,5-3.8    c4.4-4.7,9.2-8.7,15.3-10.7c-0.1-0.7,0-1.1-0.2-1.2c-1.6-1.6-4-3.6-3.2-5.6c1-2.5,4.2-2.9,6.9-2.5c5.8,0.9,10.9-1.6,16.1-3.2    c3.4-1.1,6.4-3.7,5.5-7.6c-0.8-3.8-4.7-3.7-7.9-4.3c-11.4-1.9-22.7,0.4-34,0.1c-3.6-0.1-6.9-0.6-10.5-1.8    c-8.1-2.9-14.1-7.5-17.5-15.6c-1.3-3.1-5.2-3-7.9-4.5c-2.7-1.5-6-2.5-9.1-3c-1.9-0.3-3.2-1.5-4.1-2.4c-1.6-1.5-3.6-3.9-2.5-6.3    c1.2-2.6,4.1-1.8,6.2-1.2c9.9,2.5,19.5,2.1,29.4-0.5c7.3-2,14.6-5.1,22.8-3.9c-4.5-4.3-7.2-9.9-11.1-14.5    c-3.9-4.6-10.1-3.9-17.8,1.4c-1.6,1.1-3.4,2.2-3.9,4.5c-0.5,2.6-1.4,5.4-4.5,5.7c-3.5,0.3-5.8-2.1-7.4-4.8    c-5.1-8.3-3.7-16.4,1.5-24.1c3.8-5.6,9-9.5,14.7-13.3c7.6-5.1,15.6-8.5,23.8-12.2c3.6-1.6,6.5-5.9,11-7    c10.4-2.5,20.9-1.8,31.4-0.8c13.6,1.3,26.8,4.6,40,7.3c8.6,1.8,17.1,4.2,25.6,6.5c9.6,2.6,19,3.9,28.8,0.6c2.5-0.8,3.8-2.5,5.4-4    c1.8-1.7,3.5-3.3,5.9-1.5c2.3,1.8,2.4,4,1.1,6.6c-1.7,3.4-4.3,5.7-7.6,7.4c-3.5,1.9-7.2,3.6-10.8,5.4c2.4,4.2,7.1,5.1,10,8.3    c6.4,6.9,11.2,15,18.1,21.3c1,0.9,1.4,2.2,0.5,3c-1.4,1.1-3.2,2.1-5,1.2c-2.4-1.2-4.8-2.3-8.4-2.9c7,9.2,8.9,20,14.8,28.9    c1.1,1.6,1.9,3.4,3.2,4.9c1.1,1.3,2.6,2.9,1.6,4.8c-1,1.8-2.9,2.5-5,2.5C259.7,121.1,256,120.1,252.1,118.1z"
		pathData1 := "M160.6,86.4c-0.6-3.4-4-3.7-6.2-5.2c-1.7-1.2-3.9-1.6-5.3-4.1c4.5-0.1,8.9,0.4,12.8-2.9    c-2.9-0.1-5.3-0.1-7.9,0.7c-6.1,1.9-16.5-5-20-10.4c-0.3-0.4-0.2-1-0.5-1.3c-4.6-4.4-4.7-8.7,0-13.2c0.7-0.7,0.9-2.1,0-2.7    c-1.4-0.9-1.2,0.9-1.8,1.4c-1.7,1.5-1.9,4.3-4.8,5.2c1.2-3,1.7-7.5,3.5-8.1c2.4-0.8,5.9,0.1,9.4,0.3c-2,3.5-4.1,6.7-3.7,10.7    c0.4,4,2.7,7,5.6,9.4c8.5,7,17.8,6.6,25.1,1.3c3.8-2.7,5.1-7.5,3.7-12.3c-0.3-1-0.9-2.1-0.2-2.7c1.3-1.2,1.9,0.6,2.7,1.1    c4.5,2.8,9.4,4.8,11.8,10.5c1.9,4.4,9.2,6.4,14,4.2c-6-0.2-7.2-4.4-8.7-7.9c-2-4.5-5.5-6.6-9.8-8.6c-5.8-2.6-12.4-4.1-16.3-9.9    c-0.2-0.3-0.2-1.3-0.9-0.5c-4.4,5.4-10.3,1.2-15.4,2.5c-0.4,0.1-0.8-0.1-2.3-0.3c4.5-1.8,8.1-3.4,12.4-3.9    c-4.8-4.1-8.9-0.1-12.7,0.9c-9.4,2.5-17.9,2-26.1-4c0.9-1.1,1.9-2.1,2.8-3.2c-3.1-2.2-5.5,5.9-8.6,0.1c1.7-0.9,3.2-2.3,4.4-4    c-2.6-1.7-3.6,1.4-5.4,1.9c-1.3,0.4-2.1,0.7-2.9-0.4c-1.1-1.4,0.1-2.2,1-2.7c4.1-2.6,8.7-4.2,13.6-4.2c18.9,0,37.3,3.3,55.6,7.2    c12.9,2.8,25.5,6.5,38.3,9.6c11.2,2.7,21.5-0.2,31.8-8.5c-3.2,7.9-9.5,9.8-15.2,12.3c-6,2.7-12.2,2.8-19.1-0.4    c1.3,4.4,4.4,5.3,6.4,5.5c11.4,0.9,17.9,9.1,24.4,16.5c4,4.5,8.9,8.4,12.8,14.1c-5-1.1-8.6-3.6-11.7-6.5    c-4.2-3.9-7.7-8.7-11.9-12.6c-3.6-3.4-7.8-6.2-13.5-5.8c6.6,2.9,11.1,8.3,15.9,13.2c1.1,1.1,1.1,1.1,1.1,3.2    c-2.7-0.9-3.8-4-5.5-4.6c1.7,0.5,2.6,3.9,5.6,4.5c3.4,5.1,8.2,9.4,9.4,15.9c-5.2-1.6-7.2-6.6-10.5-10.2c-1.3-1.4-2.6-2.9-3.8-4.4    c-2.4-2.9-4.6-4.2-6.1,1c4.8-1.1,6.8,2,8.8,5.5c2.9,5.2,6.5,9.7,11.6,13c1.7,1.1,3.4,2.6,4.2,4.7c2.7,7.2,7.5,13.3,11,20.2    c-5.2,0.8-14-4.1-15.8-8.7c-0.3-0.7-0.1-1.1,0.3-1.8c3.1-4.4,0.9-7.4-2.5-10.5c-3.2-2.8-6.7-5.3-9.9-8c-1.8-1.5-2.7-0.4-4-0.2    c3.7,5.9,10.8,8.5,14.6,16.1c-3.4-1.8-5.5-4.3-8.6-3c3.6,4.6,7.7,8.7,10.1,14.8c-6.2-4.1-9.3-11.4-17.2-11.9    c6.6,8.4,18.4,13.1,17.8,26.8c-2.1-1.6-3.9-3-4.9-5.1c-0.6-1.3-1.7-1.1-2.4-0.7c-1.4,0.8-0.3,1.6,0.4,2.3c2.3,2.2,3.3,5.6,6.7,6.9    c1.9,0.7,1.6,3,1.1,4.4c-2.8,8.4-3.2,17-3.4,25.8c-1.5-4.4-5.8-7.7-3.8-13.3c1.3-3.6,1.5-7.6,2.2-11.4c0.4-1.8-0.3-3.6-1.3-4.5    c-3.4-3-4.9-7.1-7-10.9c-2.1-3.8-6.4-5.4-9.8-8c-0.9-0.7-0.7-4.3-3.2-1.7c-1.9,2-0.3,3.4,1.2,4.5c5.8,4.1,10.5,9.1,13.1,16    c-4,0.8-4.7-5.4-9.2-4c1.2,3.2,3.7,4.5,6.4,5.7c4.2,1.9,5.2,4.8,4.1,9.4c-2.2,9.4-5.1,18.5-7.9,27.7c-1.1,3.5-1.2,7.2-2.2,10.8    c-1.3,4.6-0.6,9.6,4,12.8c0.8,0.6,2,1.3,1.6,2.8c-7.4-1.2-7.7-1.8-9.1-19.9c-3.9,3.9-3.2,8.7-1.6,13.3c2.2,6.3,0.5,11.9-2.7,17.1    c-4.4,7.1-4.6,14.1,0,20.7c1.6,2.3,3.6,4.2,5.3,6.4c2.6,3.3,5.6,4.2,9.4,0.5c-0.5,5.8-4.6,6.4-7.8,6.1c-3-0.3-6.3-1.9-8.4-4.4    c-4.9-5.9-10.9-11.2-11-20c-1.5,5.9,0.3,12.3,4.2,16.3c3.5,3.5,6.6,7.3,11.3,9.3c0.4,0.2,0.8,0.9,1.1,0.8    c10.6-0.7,11.1,8.2,14.1,14.9c-4.2-2.3-6-7.7-11.5-8.7c-0.6,2.2,1.1,3.1,2.3,3.9c6.4,4.7,7.1,16.3,1.3,21.7    c-0.4,0.4-1.3,0.5-1.8,0.4c-1.6-0.4-0.1-1.2,0-1.8c0.4-6-1.9-10.9-5.9-15.2c-1.6-1.7-1.6-5.3-5.6-3.7c1,1,2.1,1.9,2.1,3.4    c-3,1.9-4.5-1.8-7.2-2c-2.3-0.1-2.8-4.1-5.8-4.8c-0.6,1.4,1.3,1.7,1.1,3.1c-3.3-0.9-6.8-1.5-9.9-2.9c-5-2.1-9.9-4.5-14.7-7.2    c-8.3-4.7-11.4-17-7.6-26.9c0.6-1.4,2.1-2.6,1.3-4.6c-1.7,0.1-2.6,1-3.3,2.6c-6.8,15.1-15.5,28.9-25.8,41.8    c-6.3,8-0.8,24.8,7.3,29.2c3.8,2.1,7.2,0.8,10.5-1.1c1.2-0.7,2.5-1.7,3.3-0.2c0.5,1,0.6,2.5-1.3,3.6c-9.6,5.4-17.6,4.7-27.1-0.4    c-12.9-7-16.4-19.3-17.4-32.7c-0.3-4.1-1.8-7.5-3.4-11c-1.5-3.2-4.6-2-8.6-2.5c6.1-3,9.7-7.1,12.5-11.8c2.6-4.5,4.8-9.1,4.4-14.5    c-0.1-1.1,0.3-2.5-1.6-2.4c-1.5,0.1-1.3,1.1-1.4,2.1c-1,11-7.5,18.5-15.8,24.9c-5.9,4.6-12.4,7.9-19.3,10.7    c-5.1,2.1-6.9,6.9-8.2,12.2c-2.6-2-3.3-4.3-1.2-6.8c2-2.3,2.6-5.2,4.7-8.1c-5.5,0.3-5.8,5.4-9.1,7.7c-5.3-7.7-14.7-6.6-24.1-9    c4.5-0.9,7.6-1.4,10.7-2c-5.4-0.2-10.8-3.3-16,1.1c-0.4-2.7,2.1-2.9,1.9-4.8c-3.6,1-6.5,3-8.8,5.7c-0.8,1-1.4,3.4-3.1,2.1    c-1.3-1-2.4-2.7-1.1-5.1c1.3-2.4,3.4-3.3,5.3-4.7c1.9-1.4,5.3-1.5,5-6.3c-1.4,3.5-4,3.3-6.1,4.3c-2.6-5.3-3.7-11.1-4.8-16.9    c-1.8-9.7,6.3-23.1,15.8-26.4c2.7-1,5.7-1.4,8.2-2.7c3.4-1.8,8.3-2.6,6.8-8.8c3.2,3.4,2.6,6.7-1.2,8.9c-3.8,2.2-7.6,4.3-11.2,6.7    c-7.1,4.8-9,10.1-4.5,17.5c4,6.5,9.3,11.6,17.8,11.8c4.7,0.1,8.2-1.8,9.6-6c2.9-8.3,6.6-15.3,16.8-16.2c1.8-0.1,3.6-1.5,5-4    c-7.4,2.9-14.4,4.9-21.8,5.7c7.7-3.5,15.7-6.5,21.7-12.7c-4.7,1.9-9.1,4.2-13.5,7c7.1-6.6,6.8-17.6,14.7-23.7    c2.4-1.8-0.5-5.5,2.8-8c4.4-3.4,8.6-6.6,14.2-6.1c2.5,0.2,3.4-1.8,5.3-2.4c5.1-1.6,10.5-1.7,15.5-3.6c1.7-0.6,2.9-1.3,4-3.3    c-9.3,1.5-18.1,3.5-26.4,7.4c2.4-2.8,6.3-3.3,8.6-7.9c-9.9,3.9-18.2,9.3-27.8,11.6c10.5-5.4,17.4-16.5,29.8-18.6    c2.2-0.4,4.7-1.6,6.3,1.5c1,1.9,3.3,1.3,5,0.8c4.5-1.3,8.9-2.7,13.4-4.1c-2.9-1.9-2.9-1.9-16.7,2.8c5.3-4.5,9.5-8.1,13.9-11.9    c-2.7-1.1-4.3,0.4-5.7,1.7c-5,4.8-11.4,4.9-17.5,5.8c-2.8,0.4-4-1.7-5.9-3.9c12.8-0.2,25.7,0.5,30.6-16.3    c0.7,5.1,4.3,8.5,0.6,12.4c-0.2,0.2,0.1,1.5,0.5,1.7c0.9,0.4,1.5-0.3,2-1c2.2-3,2-6.3,1-9.5c-0.3-1.1-0.9-2.1-0.1-3.4    c2.8,1.1,5.5,2.3,8.3,3.4c-1.3,2.8-2.6,5.6-4,8.3c-0.3,0.6-0.8,0.9-0.1,1.6c0.6,0.6,0.9,0.5,1.6,0.1c3.6-2.3,5.2-6,6.6-9.7    c1.3-3.3-2-4.4-3.9-6.4c-2.6-2.8-3.1,1.4-4.5,0.7c-4.8-2.5-10.6-2.7-15.1-6c-1.2-0.9-3-2.6-4-1.3c-1.8,2.1,1.9,2.3,2.5,4.5    c-5.9-0.6-11.4-0.9-16.9-1.7c-2.8-0.4-5.6-0.7-8.3-0.2c-10.4,1.9-20.5,1.4-29.7-4.3c-3.6-2.3-6.8-5.4-8.8-9.6    c2.7-0.7,5.2-0.3,6.8,1.3c3.1,3.3,7.3,4.1,11,4.2c7.9,0.2,15.8,0,23.8-0.4c8.2-0.4,16.8,0.3,24.8,3.7c6.6,2.8,13.1,5.7,20.9,2.3    c8-3.5,10.4-12.1,4.5-19c-4.6-5.3-10.8-7.5-17.5-8.4c-9.7-1.3-19.6-1.7-28.8-4.3c-6.2-1.7-12.9-4.2-16.2-11.6    C103,66.7,96.8,62.5,87.9,64c-7.3,1.2-13.2,4-16.9,10.7c-0.6,1.1-0.3,3.3-2.3,3.2c-1.6-0.2-2-2.1-2.5-3.4    C63,66.3,65.6,59,70.9,52.9c3.2-3.8,6.4-7.9,11.4-10c1.8-0.8,2.2-0.8,3.1,0.9c3.5,6.3,8.3,10.4,16.3,9.3c7.3-1,14.2,1.3,21.1,3.1    c2.9,0.8,4.7,4.3,6.2,7.1c4.9,9.7,14.2,14.3,22.9,19.5c2.6,1.5,5.3,3.9,8.8,3.5c7.5,0.6,13.9,5.2,21.3,6    C174.6,91.3,168.2,86.5,160.6,86.4z"
		pathData2 := "M54.1,93c5.8,1.7,11.5,2.5,17.3,3c7.4,0.7,13.7-3.1,20.6-4.7c10.7-2.3,21-2.7,31.9-0.6    c9.7,1.8,19.7,1.8,29.4,4.3c3.4,0.9,6.6,2.1,9.4,4.8c-10.1-2.7-20.1-5.2-30.6-3.9c10.5,1.2,20.7,3.7,30.5,7.9    c-5.8,1.3-10.8-2.1-16.2-3c-8.3-1.3-16.6-1.1-24.7-0.5c-10.7,0.7-21.3,1.4-32,1.9c-10,0.5-19.4-1.6-28.8-4.3    C58.4,97.1,55.5,96.1,54.1,93z"
		pathData3 := "M101.9,32.9c1.3-0.4,2.2,0.6,3.2,1.1c6.4,3.1,12.7,6.3,19,9.5c2,1,4.1,1.8,2.3,4.7c-1.4,2.3-1.9,6.3-5.6,4.6    c-7.2-3.2-14.7-2.8-22.1-2.9c-4-0.1-10.3-5.9-10.6-9.5c-0.1-1.2,0.7-1.9,1.5-2.4C93.3,35.5,97.7,34.5,101.9,32.9z"
		pathData4 := "M89.6,106.3c11.2,1.1,22.2-1.6,33.3-2.5c9.4-0.8,18.5-0.9,27.6,1.3c2.8,0.7,5.4,2.1,7.4,4.4    c1.3,1.5,1.5,2.9,0.6,4.5c-0.9,1.6-2.6,1.4-4,0.9c-5.1-1.7-10.1-4.1-15.4-5.3c-7.8-1.7-15.4-2.7-23.5-1.8    C107.1,108.6,98.2,109,89.6,106.3z"
		pathData5 := "M159.3,50.4c6.1,2.1,8.5,7.2,6.3,12.5c-1.3,3.2-9.6,6.6-13.2,4.8c-4.1-2.1-9.2-3.3-10.6-9.1    c-0.9-4.1,0.7-9.4,4.3-11.2c1.5-0.8,3.1-0.7,4.5,0.8c1.8,1.9,4.5,3.6,4.1,6.3c-0.3,2-2.8,1.9-4.6,1.2c-1-0.4-1.6-3.2-3.2-0.5    c-1,1.8-1.6,3.4,0,5.4c2.8,3.4,7.3,4.6,10.9,2.5c4.1-2.4,3.5-6.2,2.6-9.9C160.2,52.3,160.3,51.1,159.3,50.4z"
		pathData6 := "M170.9,131.9c-1.1,2.4-1.2,4.6-1.8,6.7c-0.6,2.2-2.1,1.6-3.1,1.3c-1.6-0.4-0.1-1.7,0.1-2.2c1.2-2.5,1.8-5.1,1.9-7.9    c0.2-4.8,5.3-5,7.8-7.7c2.1-2.2,3.9-5.6,3.2-8.6c-0.2-1.1,0.4-1.5,0.9-2.1c1.1-1,1.8-0.1,2.5,0.6c1,0.9,2.1,2.1,1.3,3.4    c-0.9,1.4-1.5,2.8-2.2,4.3c-2.2,4.4-1.6,6.6,3,8.1c1.4,0.5,2.7,1,4.2,1.6c-4.5,3.6-6.7-1.7-9.9-2.6c-1.1-0.3-2.1-2.1-3.2-0.7    c-0.8,1-2.2,2-1.2,4.1c1,2.2,2.1,4.3,3.5,6.2c0.9,1.2,1.4,2.2,0.2,3.7C175,138.2,172.7,135.4,170.9,131.9z"
		pathData7 := "M232,151.2c-0.4,2,1.3,4.8-2.1,5.7c-1.7,0.5-8.4-3.9-8.8-5.5c-0.2-1,0.4-1.5,1.1-2c1.5-1.1,1.9,0.5,2.8,1.1    c1.4,0.9,1.6,4.6,3.4,3.1c1.5-1.4,1.5-4.8-0.4-7.1c-1.6-2-3.2-4.1-5-6c-0.9-1-1.4-2.1-0.6-3c0.8-1,2.2-0.7,2.9,0.4    C227.9,142.1,233.4,144.9,232,151.2z"
		pathData8 := "M217.3,169.7c1.8-1.9,2-5,4.8-5.7c0.8,2.5-1.8,9.3-3.9,9.8c-1.5,0.3-3.4,0.3-4.7-0.4c-1-0.5-0.4-2.2,0.1-3.5    c1.5-4.2-0.3-8.4-0.5-12.6c0-1.1-0.9-3,1-3.2c1.5-0.1,1.5,1.7,2.1,2.8C218,161.1,216.2,165.4,217.3,169.7z"
		pathData9 := "M199,172.3c4.2,1.3,2.5,4.7,3.2,7.3c2.1-1.8,2.3-4.6,3.6-7.1c1.5,4.3-1,11.1-4.3,12.3c-1.9,0.7-2.6-0.2-2.6-2.2    C199.1,179.3,199,176,199,172.3z"
		pathData10 := "M228.2,164.9c1.9-2.2,2.5-4.7,3.5-7.2c1.9,3.9-1,11.2-4.5,12.2c-1.8,0.5-2.3-0.4-2.3-2.1c0.1-3.4,0-6.9,0-10.5    C229.1,158.5,227.7,161.8,228.2,164.9z"
		pathData11 := "M160.6,86.4c3.1-0.9,6.2-0.6,9.1,0.7c5.6,2.4,11.4,4.1,17.7,5.4c-3.8,2.3-16.5-0.8-26.7-6.2    C160.7,86.3,160.6,86.4,160.6,86.4z"
		pathData12 := "M199.6,155.8c0.6-4.2,0.4-7.8,4.3-10.6c-1.2,5.3-1.4,9.8-4.1,13.7c-4.1-2.2-3.8-7.1-6.4-10.1c-0.7-0.8,0.4-0.7,0.7-0.7    c0.8-0.1,1.5-0.1,1.9,1C196.9,151,198.1,152.9,199.6,155.8z"
		pathData13 := "M156.7,171c5.2,0.1,3.8,5.3,7,6.8c0.5-3.3,0.2-6.9,4.3-8.5c-0.7,2.2-1.6,4.4-1.9,6.6c-0.3,2.3-0.5,5.4-2.7,5.5    c-2.1,0.1-2.8-3.2-3.8-5.2C158.9,174.6,158.6,172.7,156.7,171z"
		pathData14 := "M200,191.6c-0.9,4.1-1.4,7.2-2.4,10.2c-0.3,1-0.7,2.9-2.1,2.9c-0.9,0-1.9-1.4-2.4-2.3c-1.3-2.6-2.5-5.4-3.8-8.3    c4.5,0.1,3.2,5.2,6.4,6.7C196.4,197.6,195.6,193.7,200,191.6z"
		pathData15 := "M149.2,207.8c4.6,0.6,3.3,5.7,6.6,6.9c0.3-3.3,0.3-6.6,3.9-8.6c-0.7,3.1-1.2,6.3-2.1,9.3c-0.2,0.8-0.3,2.4-1.8,2.5    c-1.5,0.1-2.3-0.7-2.9-2.1C151.9,213.3,150.6,210.8,149.2,207.8z"
		pathData16 := "M125.8,198.9c0.3-3.4,0.3-6.7,4-8.7c-0.8,3.3-1.4,6.5-2.4,9.7c-0.2,0.6-0.2,2-1.6,2.1c-1.4,0.1-2.3-0.6-2.8-1.6    c-1.3-2.7-2.4-5.4-3.7-8.3C123.8,192.1,122.6,197.5,125.8,198.9z"
		pathData17 := "M213.9,121.2c2.4-0.5,2.7,2,4.7,1.7c1.2-2.7-0.3-4.7-1.9-6.8c-1-1.2-1.3-3.2,0-3.9c1.3-0.7,2,0.9,2.5,2.2    c0.6,1.5,1.4,2.9,2.1,4.3c1.6,2.8,0,5.1-2.1,6C216.7,125.9,214.8,124.3,213.9,121.2z"
		pathData18 := "M191.3,114.9c5.5,3.9,12.3,2.7,17.9,5.4C203.7,122.9,193.6,120.2,191.3,114.9z"
		pathData19 := "M184,104c0,0.7,0.2,1.5,0,2c-0.7,1.3-1,3.9-2.5,3.6c-1.7-0.3-1.6-2.2-0.8-4c0.8-2,0.1-4-2.5-4.6c-0.9-0.2-2-0.6-1.7-1.7    c0.4-1.2,1.7-1.5,2.9-1.2C182,98.6,184,101.2,184,104z"
		pathData20 := "M200.1,125.8c-6.2,6.5,0.3,7.1,3.4,9c1.4,0.9,3.4,1,4.4,3.1c-5,0.1-9.3-2.2-12.2-5.2C192.8,129.8,196.8,127.7,200.1,125.8    z"
		pathData21 := "M212.3,139.1c3.2,2.4,5.2,5.1,4.6,8.9c-0.3,1.6,1.2,4.4-0.9,4.7c-1.7,0.3-3.3-2-4.6-5c1.7,1,2.5,1.5,3.2,2    C216.7,145.5,210.6,143.4,212.3,139.1z"
		pathData22 := "M225.2,257.1c3.3,1.9,5.5,4.4,7.7,7.1c1.3,1.7,0.5,2.9-0.4,3.5c-1.2,0.8-2,0.3-2.5-1.6C229,263,224.6,261.7,225.2,257.1z"
		pathData23 := "M212,197.1c1.1,0.7,1.9,1.2,2.7,1.7c0.8-2.4-0.3-4.3-1.6-6c-0.9-1.1-2-2.5-0.8-3.4c1.6-1.2,2.6,0.9,3.3,1.9    c2.1,2.9,1.3,6.3,1.2,9.5c0,1.4-1.2,1.4-2,0.4C213.9,200.2,212.1,199.5,212,197.1z"
		pathData24 := "M137.3,259.2c0.8,0.5,1.6,1,2.3,1.5c2-4.2-4-6.3-2.4-10.6c3.2,2.1,5.2,4.8,4.7,8.8c-0.2,1.3,1,3.3-0.9,3.9    C138.6,263.7,138.2,261.1,137.3,259.2z"
		pathData25 := "M192.3,214.2c3.2,2.2,5.2,4.9,4.6,8.9c-0.2,1.3,1,3.3-1,3.8c-2.5,0.6-2.9-1.8-4.8-3.1c2.2-0.7,2.4,0.8,3.4,0.8    C196.8,220.5,190.4,218.3,192.3,214.2z"
		pathData26 := "M128.5,237.9c2.2,3,0.7,5.9,3,7.7c0.7-1.1,1.4-2.3,2.4-4.1c2.1,2.8,0.4,4.8-0.7,6.7c-0.4,0.7-1.2,1.2-2.1,0.5    C128,246.1,127.5,242.8,128.5,237.9z"
		pathData27 := "M188.2,157.1c1.6,3.6,0.2,5.6-1.3,7.4c-0.7,0.9-1.4,0.2-2.1-0.3c-2.3-1.7-3-4.9-2-9.5c2.5,1.8,0.1,5.4,2.8,6.8    C186.2,160.4,187,159.2,188.2,157.1z"
		pathData28 := "M113.4,223.2c0.7,0.5,1.4,1,2.2,1.5c2.1-4.2-3.9-6.3-2.3-10.4c3.1,1.8,4.9,4.3,4.8,8c-0.1,1.6,0.5,3.7-1.3,4.4    C114.4,227.6,114.2,224.9,113.4,223.2z"
		pathData29 := "M111.6,205.5c0.7-1.2,1.4-2.4,2.3-4.1c2.1,2.9,0.5,4.8-0.7,6.8c-0.4,0.7-1.2,1.1-2.1,0.4c-3.1-2.6-3.9-5.8-2.7-10    C111.2,200.5,108.9,204,111.6,205.5z"
		pathData30 := "M178.8,200c2.1,3,0,6.1,2.9,7.3c0.6-0.9,1.4-2.2,2.3-3.7c2.2,2.9,0.4,4.8-0.9,6.7c-0.8,1.2-1.7,0.6-2.4-0.2    C178,207.5,177.4,204.5,178.8,200z"
		pathData31 := "M146.8,176.9c1.9,3.2,0.3,6.1,2.7,7.7c0.7-1.2,1.6-2.4,2.5-3.9c2.2,2.9,0.4,4.8-0.9,6.7c-0.8,1.2-1.7,0.6-2.4-0.2    C146,184.5,145.4,181.5,146.8,176.9z"
		pathData32 := "M182.3,117.4c5.4,2.1,8.8,4.4,11.1,8.3c0.6,1.1,0.7,1.7-0.1,2.5c-1.3,1.3-1.9-0.3-2.1-0.8    C189.4,123.7,185.7,121.7,182.3,117.4z"
		pathData33 := "M200.1,55.7c5.3-0.6,9.3,3.9,14.5,3.6C208.8,62.2,204.5,58.7,200.1,55.7z"
		pathData34 := "M109,178c3.9,0.5,3.8,4.6,6.8,5.9c0-2.1,0.9-3.6,2.9-4.5c1.6,3.2-0.8,5.1-2.8,6.2c-2.3,1.2-3-1.4-3.9-3    C111.1,181.1,109.2,180.5,109,178z"
		pathData35 := "M153.8,162.7c0.1-2,0.8-3.7,3-4.4c1.3,3.3-1,5.4-3.2,6.3c-2.1,0.9-2.9-1.7-3.9-3.4c-0.8-1.3-2.6-2-2.7-4.2    C150.9,157.2,150.7,161.5,153.8,162.7z"
		pathData36 := "M192.7,135.2c1.2,3.6-1.2,5.6-3.3,6.5c-1.9,0.8-2.9-1.9-3.8-3.6c-0.7-1.3-2.9-1.7-2.3-4.1c3.6,0.4,3.5,4.5,6.4,5.8    C189.8,137.8,190.6,136.2,192.7,135.2z"
		pathData37 := "M185.7,185.5c0.2-1.7,0.9-3.4,2.9-4.2c1.5,3.2-0.9,5.2-2.9,6.2c-2.2,1.1-3-1.4-3.9-3.1c-0.7-1.4-2.7-1.9-2.8-4.2    C183.1,179.9,182.5,185,185.7,185.5z"
		pathData38 := "M209.7,65.3c-2.5,2.9-7.8,3.8-10.2,1.7c-0.9-0.8-2.4-1.6-0.9-3.1c1.4-1.5,1.9-0.8,2.6,0.8c0.8,1.8,2.8,1.9,4.3,1.1    C207,64.9,208.3,64.9,209.7,65.3z"
		pathData39 := "M158.3,155.1c5.1-0.4,9,3,13.2,5.3C166.2,161.2,162.7,156.9,158.3,155.1z"
		pathData40 := "M133.8,181.6c-5.4,0.7-9-3.5-13.4-5.4C125.6,175.9,129.6,179.1,133.8,181.6z"
		pathData41 := "M144.5,217.7c1.1,3.1,0.1,5.3-1.8,6.9c-1.6,1.4-2.4-1-3.5-1.6c-2.9-1.5,1.3-3.1,0.2-4.7c0.7,1.2,1.4,2.4,2.1,3.6    C142.5,220.5,142.5,218.8,144.5,217.7z"
		pathData42 := "M218.2,108.1c3.6,1.1,7.4,0.6,10.9,2.5c0.9,0.5,2.4,0.6,1.9,2.2c-0.7,2-1.9,0.1-2.5-0.1c-3.6-0.9-7-2.2-10.5-3.4    C218,108.9,218.1,108.5,218.2,108.1z"
		pathData43 := "M205.4,127.5c4.3-0.3,7.8,1.8,12.6,4.2C212.1,132.5,208.7,130.1,205.4,127.5z"
		pathData44 := "M203.1,106.4c6.1-3.7,10.3-1.2,14.5,0.5C213.1,108.1,208.9,105.6,203.1,106.4z"
		pathData45 := "M200.7,110.4c2.6,2.7,5.1,4.7,8.7,3.7c-2.1,2.1-4.2,2.4-6.9,1.1C200.4,114.2,199.5,113.3,200.7,110.4z"
		pathData46 := "M198.2,165.8c1.4-3,5.2-3,5.9-6.1c0.2-0.8,1-0.9,1.5-0.5c0.4,0.4,0.6,1.4,0.4,1.8C204,163.5,202.5,166.8,198.2,165.8z"
		pathData47 := "M195.5,89.4c-8.1,0.9-8.3,0.8-10.5-3.5C188.8,86.9,191.8,89.4,195.5,89.4z"
		pathData48 := "M193.8,108.3c-0.2,0.4-0.3,0.6-0.4,0.6c-2.7,0.5-5.2,0-7.1-2.1c-0.6-0.6-0.4-1.5,0.3-1.6    C189.5,105,191.6,106.8,193.8,108.3z"
		pathData49 := "M162.5,147.2c0.6,1.1,1.2,2.2,2.1,3.7c0.6-2.5,1.5-4.5,4.8-5.1c-1.5,2.9-2.5,5.7-5.2,6C162.5,152,162.1,149.5,162.5,147.2    z"
		pathData50 := "M239.3,71.7c-4.5,0-5.8-5.3-9.8-6c5.1-0.9,6.7,3.8,9.7,6.2C239.2,71.8,239.3,71.7,239.3,71.7z"
		pathData51 := "M148.7,195.4c1.3,3-1,4.4-2.7,6.1c-0.5,0.5-1.2,0.9-1.7,0.1c-0.3-0.5-0.2-1.2-0.2-1.9    C146.9,199.6,145.9,195.4,148.7,195.4z"
		pathData52 := "M201,102.7c-0.3,1.5-1.4,1.5-1.9,1.2c-1.6-1.2-4.4-0.2-5-2.9c-0.2-0.9,0.8-1.2,1.3-0.9C197.1,101.4,199.6,101.1,201,102.7    z"
		pathData53 := "M226.1,174.6c-0.9,0.7,2.1,2.2-0.5,2.4c-2.8,0.2-1.2-2.8-2.5-3.9c-0.5-0.4-0.3-2,1.2-2C227.3,171,225.4,173.5,226.1,174.6    z"
		pathData54 := "M216.9,96.1c-2.6,1.4-4.9,1-7.2-1.1C212.7,92.7,214.7,95.3,216.9,96.1z"
		pathData55 := "M217.2,85.3c-2.3,2.2-4.3,2-6.3,1c-0.5-0.3-1-0.8-0.3-1.5c1.3-1.4,2.5-1,3.9-0.1C215.1,85.2,216,85.1,217.2,85.3z"
		pathData56 := "M192.3,83.4c2.4-3.2,4.5,0.3,7.6-0.2C196.1,85.5,196.1,85.5,192.3,83.4z"
		pathData57 := "M172.9,191.4c0.2,2.7-1.1,4.3-3.5,4.4c-1.2,0-2-1.7-1.3-3.2c0.2-0.5,0.8-2.4,1.7-0.6C171.9,195.6,171.6,191.2,172.9,191.4    z"
		pathData58 := "M104,220.7c0,3.4-1.1,5.1-3.6,5c-1.2-0.1-1.9-1.8-1.2-3.2c0.2-0.5,0.5-1.6,1.6-0.9C102.4,222.6,102.7,222.6,104,220.7z"
		pathData59 := "M158,235.7c0,3.4-1.1,5.1-3.6,5c-1.2-0.1-1.9-1.8-1.2-3.2c0.2-0.5,0.5-1.6,1.6-0.9C156.4,237.6,156.7,237.6,158,235.7z"
		pathData60 := "M224.9,68.4c0.2,2.7-1.1,4.3-3.4,4.4c-1.1,0-2-1.7-1.3-3.2c0.2-0.5,0.8-2.4,1.7-0.6C223.8,72.6,223.6,68.2,224.9,68.4z"
		pathData61 := "M205.2,47.2c0.8-0.1,1.2-0.3,1.4-0.2c2.3,0.8,4.7,1.6,8,2.7C210.5,50,207.2,51.2,205.2,47.2z"
		pathData62 := "M228.1,89.3c-0.8,0.7-1.5,1.7-1.9,1.6c-1.6-0.5-3.5-0.8-4.2-2.8c-0.2-0.8,0.4-1.1,1.2-1.1    C225.1,87.1,226.4,88.2,228.1,89.3z"
		pathData63 := "M150.8,192.3c0.9-0.1,1.4-0.3,1.5-0.2c1.5,1.8,5-0.3,5.6,2.9c0,0.2-0.3,0.8-0.4,0.8C155.2,195.3,153,194.3,150.8,192.3z"
		pathData64 := "M204.8,79.5c-1.6,2.1-3.6-1.9-5.2,0.4c-0.5-0.4-0.9-0.9-1.4-1.3c0.5-0.5,1.1-1.4,1.7-1.4C202,76.9,203.5,78,204.8,79.5z"
		pathData65 := "M227.6,80.8c0.9-0.7,2.1-1.7,2.3,0.3c0.1,1.1-1.1,1.9-2.5,1.8c-1.3-0.1-2.5-0.8-2.3-2C225.6,78.9,226.7,80.6,227.6,80.8z"
		pathData66 := "M171.7,69.9c1.5-2.7,0.8-5,3.1-5.8c0.3-0.1,1.1,0.3,1.1,0.6C176.3,67.1,174.6,68,171.7,69.9z"
		pathData67 := "M216.3,98.4c3-1.5,4.3,1.1,6.2,1.9C219.9,100.9,219.9,100.9,216.3,98.4z"
		pathData68 := "M207,43.6c-2.7,1.1-4.4,0.1-5.7-2.1C204.1,39.8,204.8,43.2,207,43.6z"
		pathData69 := "M169.5,73c-0.9-0.1-1.6-0.3-1.4-1.1c0.3-1.2,1.1-1.9,2.4-1.9c0.9,0,1.8,0.3,1.4,1.3C171.5,72.4,170.6,73.2,169.5,73z"
		pathData70 := "M222.8,65.3c-0.9,1.8-2.6,2.2-4.2,2.6c-0.1,0-0.6-1.1-0.5-1.1C219.4,65.8,220.7,64.5,222.8,65.3z"
		pathData71 := "M101.8,218.3c-0.9,1.8-2.6,2.2-4.2,2.6c-0.1,0-0.6-1.1-0.5-1.1C98.4,218.8,99.7,217.5,101.8,218.3z"
		pathData72 := "M155.9,233.2c-1.1,2.3-3.1,1.9-4.6,2.6C152.1,233.4,153.6,232.7,155.9,233.2z"
		pathData73 := "M171.4,188.3c-2.1,1.6-3.4,2.5-5.9,2.6C167.1,188.4,168.6,187.7,171.4,188.3z"
		pathData74 := "M208.3,179.5c0.4,1.8,1.6,4-0.1,5.2c-1,0.7-1.3-0.5-0.8-1.7C207.9,182,208,180.7,208.3,179.5z"
		pathData75 := "M138.3,167.3c0.9,0.1,0.7,0.8,0.7,1.3c0.1,1.1-0.6,1.4-1.5,1.3c-0.6-0.1-1.4,0-1.4-1C136.1,167.6,137.7,168.2,138.3,167.3    z"
		pathData76 := "M238.2,118.3c2.1-0.8,3.3-0.1,4.1,2.9C240.5,120,239.4,119.2,238.2,118.3z"
		pathData77 := "M216,78.1c0.2-1.3,1.1-1.1,1.8-1.1c0.3,0,1.1-0.1,0.7,0.7c-0.3,0.7-0.9,1.4-1.8,1.2C216.4,78.8,216.2,78.3,216,78.1z"
		pathData78 := "M162.3,75.8c1.6-2.4,2.5-3.2,4.1-2.6C165.9,74.6,164.7,74.6,162.3,75.8z"
		pathData79 := "M141.6,195.3c0.7,0.4,1.4,0.9,1.3,1.8c0,0.3-0.5,0.7-0.9,0.8c-0.9,0.2-1-0.5-1-1.2C141,196.2,140.8,195.5,141.6,195.3z"
		pathData80 := "M103.2,39.7c-2.6,0.2-4.2,1.3-5.9,2.3c1.8,2.9,5.5,1.1,7.4,3.8c-3.9,0.6-7.5,0.5-10.2-2.8c-0.9-1.1-0.2-2.2,0.6-2.4    C97.6,40,99.6,36.1,103.2,39.7z"
		canvas.Translate(20, 30)
		canvas.Scale(0.9)
		canvas.Path(pathData0, "fill:"+lineColor)
		canvas.Path(pathData1, "fill:"+hexCode)
		canvas.Path(pathData2, "fill:"+hexCode)
		canvas.Path(pathData3, "fill:"+hexCode)
		canvas.Path(pathData4, "fill:"+hexCode)
		canvas.Path(pathData5, "fill:"+hexCode)
		canvas.Path(pathData6, "fill:"+lineColor)
		canvas.Path(pathData7, "fill:"+lineColor)
		canvas.Path(pathData8, "fill:"+lineColor)
		canvas.Path(pathData9, "fill:"+lineColor)
		canvas.Path(pathData10, "fill:"+lineColor)
		canvas.Path(pathData11, "fill:"+lineColor)
		canvas.Path(pathData12, "fill:"+lineColor)
		canvas.Path(pathData13, "fill:"+lineColor)
		canvas.Path(pathData14, "fill:"+lineColor)
		canvas.Path(pathData15, "fill:"+lineColor)
		canvas.Path(pathData16, "fill:"+lineColor)
		canvas.Path(pathData17, "fill:"+lineColor)
		canvas.Path(pathData18, "fill:"+lineColor)
		canvas.Path(pathData19, "fill:"+lineColor)
		canvas.Path(pathData20, "fill:"+lineColor)
		canvas.Path(pathData21, "fill:"+lineColor)
		canvas.Path(pathData22, "fill:"+lineColor)
		canvas.Path(pathData23, "fill:"+lineColor)
		canvas.Path(pathData24, "fill:"+lineColor)
		canvas.Path(pathData25, "fill:"+lineColor)
		canvas.Path(pathData26, "fill:"+lineColor)
		canvas.Path(pathData27, "fill:"+lineColor)
		canvas.Path(pathData28, "fill:"+lineColor)
		canvas.Path(pathData29, "fill:"+lineColor)
		canvas.Path(pathData30, "fill:"+lineColor)
		canvas.Path(pathData31, "fill:"+lineColor)
		canvas.Path(pathData32, "fill:"+lineColor)
		canvas.Path(pathData33, "fill:"+lineColor)
		canvas.Path(pathData34, "fill:"+lineColor)
		canvas.Path(pathData35, "fill:"+lineColor)
		canvas.Path(pathData36, "fill:"+lineColor)
		canvas.Path(pathData37, "fill:"+lineColor)
		canvas.Path(pathData38, "fill:"+lineColor)
		canvas.Path(pathData39, "fill:"+lineColor)
		canvas.Path(pathData40, "fill:"+lineColor)
		canvas.Path(pathData41, "fill:"+lineColor)
		canvas.Path(pathData42, "fill:"+lineColor)
		canvas.Path(pathData43, "fill:"+lineColor)
		canvas.Path(pathData44, "fill:"+lineColor)
		canvas.Path(pathData45, "fill:"+lineColor)
		canvas.Path(pathData46, "fill:"+lineColor)
		canvas.Path(pathData47, "fill:"+lineColor)
		canvas.Path(pathData48, "fill:"+lineColor)
		canvas.Path(pathData49, "fill:"+lineColor)
		canvas.Path(pathData50, "fill:"+lineColor)
		canvas.Path(pathData51, "fill:"+lineColor)
		canvas.Path(pathData52, "fill:"+lineColor)
		canvas.Path(pathData53, "fill:"+lineColor)
		canvas.Path(pathData54, "fill:"+lineColor)
		canvas.Path(pathData55, "fill:"+lineColor)
		canvas.Path(pathData56, "fill:"+lineColor)
		canvas.Path(pathData57, "fill:"+lineColor)
		canvas.Path(pathData58, "fill:"+lineColor)
		canvas.Path(pathData59, "fill:"+lineColor)
		canvas.Path(pathData60, "fill:"+lineColor)
		canvas.Path(pathData61, "fill:"+lineColor)
		canvas.Path(pathData62, "fill:"+lineColor)
		canvas.Path(pathData63, "fill:"+lineColor)
		canvas.Path(pathData64, "fill:"+lineColor)
		canvas.Path(pathData65, "fill:"+lineColor)
		canvas.Path(pathData66, "fill:"+lineColor)
		canvas.Path(pathData67, "fill:"+lineColor)
		canvas.Path(pathData68, "fill:"+lineColor)
		canvas.Path(pathData69, "fill:"+lineColor)
		canvas.Path(pathData70, "fill:"+lineColor)
		canvas.Path(pathData71, "fill:"+lineColor)
		canvas.Path(pathData72, "fill:"+lineColor)
		canvas.Path(pathData73, "fill:"+lineColor)
		canvas.Path(pathData74, "fill:"+lineColor)
		canvas.Path(pathData75, "fill:"+lineColor)
		canvas.Path(pathData76, "fill:"+lineColor)
		canvas.Path(pathData77, "fill:"+lineColor)
		canvas.Path(pathData78, "fill:"+lineColor)
		canvas.Path(pathData79, "fill:"+lineColor)
		canvas.Path(pathData80, "fill:"+lineColor)
		canvas.Gend()
		canvas.Gend()
	},
}