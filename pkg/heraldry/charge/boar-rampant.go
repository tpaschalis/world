package charge

import svg "github.com/ajstarks/svgo"

var BoarRampant = Charge{
	Identifier: "boar-rampant",
	Name:       "boar rampant",
	Noun:       "boar",
	NounPlural: "boars",
	Descriptor: "rampant",
	SingleOnly: false,
	Tags: []string{
		"aggressive",
		"animal",
		"boar",
		"hunting",
	},
	Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
		pathData0 := "M240.3,281.1c3.4-2.4,5.6-4.2,7.1-6.7c1-1.7,1.1-3.4-0.8-5.2c-5.2-4.9-5-8.3,0.1-13.5c1.3-1.3,2.6-2.8,3.8-4.2    c2.1-2.5,1.2-5-0.2-7.4c-1.6-2.6-4.3-0.6-5.3,0c-4.2,2.5-8.1,0.5-12,0.1c-6.4-0.6-12-2.9-16.5-9.3c-0.2,4.1-2.2,6.4-2,9.6    c0.1,2.6,0.6,4.8,2.9,6.1c2.9,1.7-1.8,2.8,0.4,4.3c4,2,8.7,3.7,13.9,3.4c1.4-0.1,2.1,1,2.7,1.9c0.7,1.2-0.8,1.7-1.4,2.4    c-1.1,1.4-2.8,1.6-4.3,1.8c-1.8,0.3-3.2,1-3.2,2.6c0,1.9,1.1,3.6,3,4.3c0.8,0.3,2.1-0.6,2.3,1.2c0.1,1.5-0.7,2.4-1.7,3.1    c-1.3,0.8-2.7,1.3-4.3,1.1c-1.1-0.2-2,0.3-2.4,1.4c-0.4,1.2,0.4,1.9,1.2,2.6c1,0.9,3.4,1.2,2.4,3.1c-0.9,1.6-2.1,3.6-4.5,3.6    c-1.5,0-1.9,0.4-2.5,1.9c-0.9,2.2-2.2,4.2-4.4,5.7c-2.1,1.4-3.7,3.4-2.5,6.7c1.4,3.8-2.2,9.5-6.1,10.2c-2,0.4-4.3,0.6-3.6-2.8    c0.1-0.7,0-1.4,0-2.9c-2.5,2.7-4.8,4.4-7.8,5.1c-1,0.2-1.6,1.4-2.2,2.4c-2.4,3.7-6.8,4.1-10,0.8c-1.3-1.3-2.3,0.2-3.4-0.2    c-0.8-0.3-2.1,0-2.5-0.5c-0.7-0.9,0.1-2.1,0.7-2.7c4.4-4.7,7.1-11.3,14.6-12.2c2-0.2,3.1-1.2,3.9-3.6c1.9-5.7,2.1-12,6.2-16.8    c0.8-0.9,0.6-1.2-0.5-1.8c-1.2-0.7-2.4-2-3.7-1.3c-5.1,2.8-9-0.9-13.3-2.3c-1.4-0.5-2.8-1.5-4.1-1.5c-2.1,0-3-0.9-3.6-2.6    c-0.8-2.4-1.9-2.7-3.6-0.7c-0.7,0.8-1.5,2.1-2.8,1.6c-1.9-0.8,0.3-1.8,0.1-2.8c-0.2-1.2-0.5-2.2-1.2-3.3c-3-4.8-7.5-6.9-12.8-8.7    c1.5-0.7,3-1.4,4.5-2c0.9-0.3,1.8-0.7,1.6-1.7c-0.2-1.1-1.2-0.6-1.9-0.6c-1.8,0-3.6-0.1-5.3,0.3c-0.4,0.1-2.4,0-1.6,1.8    c0.4,1,1.5,2.1-0.5,2.5c-1.5,0.3-2.4-0.6-3.2-1.6c-0.3-0.4-0.2-1-0.3-1.4c-0.7-2.8-1.3-3.1-4.5-1.6c-4.7,2.2-3.4,4.3-1.4,7.4    c-2.6,0.8-4.8,0.4-6.1-1.6c-1.4-2.2-1.9-0.6-2.4,0.4c-0.5,1-0.6,2.1-0.8,3.2c-0.3,1.5-0.7,3.4-2.1,3.6c-1.5,0.2-3-1.5-3.1-2.8    c-0.6-5.3-5.9-7.9-7-12.8c-0.4-1.7-1.9-1.8-3-2.5c-2.3-1.4-4.9-2.2-7-4.2c-1.5-1.5-2-3-1.3-4.7c0.3-0.9,0.9-2.1,0.3-2.4    c-4.1-1.8-4.7-6-6.8-9.2c-1.8-2.8-5-3.3-7.8-3.9c-2.7-0.6-4.1-2.4-5.7-4.2c-1.5-1.6-0.8-2.3,0.9-2.8c0.7-0.2,1.9,0.5,2-1    c0.1-1.2-0.7-2-1.6-2.4c-2.3-1.2-3.9-3-5.2-5.3c-1-1.8-2.8-3.5,1.1-4.6c2.1-0.6,0.9-3.3,0.7-4.9c-0.1-0.8-0.8-1.8-0.1-2.4    c0.7-0.6,1.6,0,2.4,0.5c4.3,2.9,8.9,5.1,13.7,6.9c2.5,0.9,4.1,3,5,5.3c1.9,4.9,5.3,8.7,8.4,12.6c2.6,3.3,8.5,2.8,10.9-0.6    c7.2-10.6,18.6-15.3,31.8-9.7c0.4,0.2,1,0,1.9,0c-3.3-4-5.7-8.1-9.8-11.1c-1.5-1.1-4-3.5-1.7-6.3c0.2-0.2,0.3-0.4,0.4-0.6    c-1.6-0.6-3.5-0.7-4.4-2.4c-1-2.1-2.1-2.4-3.7-0.8c-0.9,0.8-1.9,0.9-2.7,0c-1.8-2.1-4.3-3.7-5.5-6.3c-0.7-1.6-1.3-2.8-3.5-1.6    c-1.3,0.7-2.4-0.7-2.2-2.2c0.2-2-0.4-3.8-1.6-5.2c-2.2-2.8-1.8-5.4-0.1-8.2c-1.6-0.7-2.8-0.4-3.8,0.9c-1.3,1.7-3.3,2.7-4.9,4    c-2,1.7-4.1,2.3-6.2,0c-1-1.1-2.2-0.9-3.2-0.6c-7.9,2.3-16,3.9-23.5,7.4c-1.2,0.6-2.3,1.2-2.7,2.4c-1.7,5.1-6.3,6.1-10.6,6.8    c-3.4,0.5-3.2-2.6-3.3-5.1c-1.8,1-2.7,2.5-3.6,4c-2.3,4.1-10,6.8-14.2,5c-0.2-0.1-0.6-0.2-0.6-0.4c-0.8-5.4-5-2.8-7.8-3.2    c-1.2-0.2-2.2-0.5-3.5-1.3c5.4-5.2,10.1-11.2,17.9-13.3c3.1-0.9,5.9-2.7,8.1-5c4.6-4.8,9.5-9,15.9-11.5    c7.2-2.8,12.7-8.4,19.4-12.2c1.6-0.9,1-1.7,0.5-2.4c-1.2-1.5-2.5-3.2-4.1-4c-4-2-4.3-5.8-3.1-8.7c1.1-2.7,0.1-4.2-1.2-5.7    c-4.2-4.8-6.7-10.5-9.1-16.3c-1.6-3.7-2.9-7.2-8-8.2c-4.1-0.9-6.7-5.2-9.4-8.5c-1.5-1.9,0.2-3.9,2.9-4.1c2-0.2,4.9,1.3,5.9-1.2    c1-2.4-1.9-3.8-3.4-4.9c-4.9-3.7-6.5-8.4-5.6-14.2c0.2-1.1-0.3-2.4,0.9-3.1c1.5-0.9,1.8,0.9,2.8,1.3c1.1,0.5,2,1.3,3.4,2.4    c-1.4-3.8-1-7.3-0.5-10.9c4.2,4.5,9.3,8.1,10.3,14.8c0.8,5.4,3.8,9.5,8.5,12.5c3.9,2.4,6.8,5.9,9.6,9.5    c5.4,7.2,9.1,15.6,16.5,21.3c1.8,1.4,2.8,0.9,4.4,0.1c2-1,3.2-3.1,5.4-4c1.6-0.6-0.4-1.1-0.5-1.7c-0.3-1.4-1.9-2.9-0.5-4.1    c1-1,2.4,0.3,3.3,0.9c4.7,3.1,8.5-0.4,12.5-1.7c0.4-0.1,0.7-0.3,1.2-0.4c-1.6-2.3-4.1-2.8-6.5-3.7c-6.1-2.1-12.7-3.6-15.5-10.5    c-0.6-1.5-1.6-0.8-2.5-0.9c-2.4-0.4-5.3-0.9-5-3.8c0.3-2.4,2-4.6,5.4-4.7c3.8-0.1,7.1,1.2,10.5,2.1c1.2,0.3,2.3,0.2,3.4,0.3    c1.6,0.1,1.5-1.5,1.9-2.4c0.5-1.1-0.7-1.1-1.4-1.5c-2.4-1.5-5.1-0.2-7.7-1.3c-4.9-2-8.4-6.3-8.4-11.5c0-1.4,0-2.8,0-4.2    c0-3.2,2-5.6,4.5-6.4c2.6-0.8,6.1-1.4,7.9,2.4c0.6,1.2,1.7,1.2,2.9,1.2c7.8,0,15.6,0.1,23.4,1.3c6.8,1.1,13.4-2.5,20.1-3    c9.4-0.8,14.8-7.1,20.7-13.1c2.7-2.8,5.6-4.7,9.8-4.5c6.2,0.3,6.2,0.3,8.2,6.5c0.7,2.3,0.6,4.6,1.3,6.9c0.7,2.7,0.2,5.5-2.3,7.8    c9-0.8,15.8,4.6,23.5,7.1c1.7,0.5,3.1,1,3.4,2.9c0.4,1.9-1.1,2.9-2.2,3.7c-6.4,4.3-12.8,8.8-20.8,9.7c-2,0.2-2.2,1-1.4,2.6    c1.8,3.6,2.5,7.5,2.3,11.5c-0.2,6,1.9,11.7,2.1,17.7c0,1.1,0.1,2.3,0.1,3.4c0,12.9,0,25.8,0,38.7c0,16.4,2.6,32.7,1.4,49    c-0.4,6.1,0.2,12,0.1,18c0,3.3,0.8,6.4,1.7,9.5c1.4,5.2,6.2,5.9,10.3,7.3c2.8,1,5.8,1.4,8.7-0.2c1.1-0.6,2.1-0.8,0.8-2.6    c-1.5-1.9-2.7-4.1-3.5-6.4c-1.5-4,0.4-10.8,3.9-12.8c3.8-2.2,10.4-0.6,13.2,3c4,5.4,3.7,13.4-1.4,18c-1.8,1.6-3.1,2.9-0.3,4.9    c5.9,4.4,4.5,11.7,6.5,17.6c0.7,2.1,1.8,4.9,0.5,6.6c-3.9,5-5.1,12.5-12.1,14.5C248.1,281.8,244.6,281.1,240.3,281.1z"
		pathData1 := "M204.7,250.8c3.1-0.8,4.6-3.9,7.6-5.7c-0.3,4.7-3.5,7.3-6,9.8c-1.7,1.7-0.6-1.3-1.3-1.9    c-1.2,0.5-1.3,2.5-3.1,2.4c-0.5,0-1.9,0.9-1.3,2.1c0.4,0.9,1.4,1,2.3,0.5c3.4-1.8,7.3-2.9,9.1-6.8c0.5-1,0.9-1.9,2.1-1.3    c0.9,0.4,2,1.2,1.9,2.1c-0.1,0.8-0.2,2.7-2.1,1.6c-0.1-0.1-0.8,0.4-0.7,0.6c0.2,1.2,1.1,2,2.2,2.4c3.2,1.1,6.5,2.2,9.8,3.1    c1.9,0.5,3.8,1.4,6.5,0.4c-3.4,3.5-6.4,3.2-9.9,1.2c1.4,1.7,2,3.2,1.4,5.5c-0.7,2.5,1.8,5.2,5,6.5c-4.9,2.9-7.7,2.6-11.9-1    c1.5,4.2,4.2,7.6,7.1,10.9c-3.2,3.5-4.3,3.4-6.4-0.7c-0.6-1.2-0.8-2.6-2.1-3.4c-0.5,3.7,4.6,7.3,0.3,11.2    c-1.5-0.9,0.2-1.9-0.2-2.7c-0.7-0.3-1.3,0-1.3,0.5c-1.1,5.4-6.5,9.9-3.7,16.3c0.5,1.1-3.4,4.9-4.7,5c-0.3,0-0.4,0.2-0.5-0.2    c-0.1-0.6-0.2-1.2,0.3-1.8c1.1-1.2,3.1-1.1,4.1-2.8c-1.8-0.8-3.3-0.1-5,0.5c-0.1-2.2-0.2-4.2-0.3-6.1c-2.3,6.8-6.6,10.4-14,9.4    c0.3,1,3,1.1,1.5,2.8c-1.6,1.9-3.2,4.2-6.6,2.6c3-2.4,3.1-6.8,6.6-8.7c0.4-0.2,0.7-0.9,0.2-1.3c-0.3-0.3-1.1-0.4-1.4-0.3    c-1.3,0.8-2.3,1.8-3,3.2c-1.5,3-4.3,4.3-8.4,5.3c2-2.3,4.2-3.3,5.1-5.4c0.2,0.6,0.3,1.2,0.5,1.8c1-1,1.4-1.8,0.9-3.6    c-0.5-1.7,2.8-2.8,4.2-2.8c7.1,0.1,9.3-4.7,10.4-10.2c1.2-6.2,4.3-11.5,7.1-16.9c0.3-0.6,1.3-0.9,0.3-1.8    c-0.9-0.7-1.1,0.2-1.4,0.6c-1.1,1.6-1.7,0.6-2.9-0.1c-3-1.8-5.2-4.5-7.7-7.1c-0.9,2.7,0.6,4.2,1.8,5.9c-4,1.1-4.6,0.5-5.6-5.9    c-2.2,0.9-0.4,2.2-0.4,3.2c-0.1,1.2,1.6,2.1,0.6,3.5c-3.3-1.7-4.5-4.8-5.3-8.1c-3.1,2.1,0.2,3.8,0.1,5.6c-2.2-1.1-2.2-1.1-5.7-7.8    c-2.9,1.6,0.9,4-0.7,5.7c-2.2-2.5-2.5-3.3-3-8.8c-1.4,2.1-2.8,4.1-4.4,6.4c-1.3-2.4-1.4-5.1-0.5-6.7c1.7-3.1,0.1-4.9-1.7-7.8    c0.2,3.1,1,5.7-1.5,7.8c0.9-8-6.4-1.4-8.2-4.8c6.3-1.3,10.4-5.3,10.3-10.3c-0.9,1.4-1.9,2.6-3.2,3.4c-2.2-8.8,0.8-16.4,5.8-23.1    c2.2-2.9,3.7-6.2,5.9-9.1c0.2-0.2,0.1-0.9-0.1-1c-0.7-0.4-1.3,0-1.6,0.6c-1.5,4-4.2,0.6-4.9,0.1c-3.7-3.3-7-7.1-10.4-10.7    c-0.8-0.8-1.8-1.6-1-2.9c0.2-0.4,1-0.9,1.1-0.8c1,0.9,2.6,0.3,3.4,2c1.5,3,3,6,5.9,8.2c2.8,2.1,4.2,1.1,5.8-0.9    c1.9,0.5-1.3,2.5,0.7,2.6c1.9,0.2,2-1.7,2-2.9c0-3.8,2.8-7.3,1.2-11.5c-1.5,1.8-1.2,4-1.8,5.9c-0.3,0.9,0.1,2-0.7,2.8    c-0.5,0.5,0.5,2-1.2,1.8c0.7-2-2.2-0.8-2-2.4c0.1-0.3-0.1-1.3,0.5-0.8c0.7,0.6,1.2,1.7,1.6,0.1c0.2-0.9,1-2.6-1-3.1    c-2-0.5-2.1-2-2.1-3.8c0-1.4,0.7-2.4,1-3.6c1.1,0.4,1.9,0.9,0.9,2.2c-0.3,0.4-0.3,1-0.6,1.3c-1.7,2,0.9,1.2,1.2,1.5    c1.3,1.2,0.6-0.6,0.6-0.7c-0.2-1.2,0.8-1.8,1.4-2.6c-0.7-0.5,3.1-2.7-0.8-2.1c-0.7,0.1-0.6-0.9-0.1-1.4c3.6-4,0.8-9.2,2.4-13.6    c-1,3.2-3,5.7-4.6,8.6c1.8,0.5,1.6-2.2,2.4-1.4c2.1,2-0.9,3.5-0.8,5.5c-0.9-2-3.3-1.7-4.5-3.5c1.6,2.9,2.3,5.6-0.6,7.7    c-1.8-4.3-3.4-9.1-3.2-14.5c-1.6,1.6-1.9,2.8-1.3,4.8c1.5,5,2.4,9.9-0.7,14.6c-3.2-7.4-9.7-13.5-8-22.8c-1.9,2.5-1.9,4.3-0.1,10.4    c-1.7-1.2-1.7-1.2-2.9-6.4c-0.4-1.7-0.8-2.4-1.4-2c-1,0.6-1.3,1.6-0.6,2.6c0.6,0.8,0.3,1.7,0.5,2.6c0.2,0.9,1.8,2,0.1,2.5    c-1.3,0.4-2.1-1.3-2.3-2.4c-0.3-2.3-0.1-4.7-0.5-7c-2.8,2.3,1.6,5.7-1.4,7.8c-3.2-4.7-0.7-10-1.4-15.6c-3.3,3.4-0.3,7.5-2,10.9    c-1.9-1.9-3.6-3.7-1.6-6.7c1.6-2.4-0.8-5.2,0.7-7.7c-5.2,4.2-1.8,10.5-4.2,16c-0.7-4.1-3.5-7.2-1.9-11.2c0.2-0.5,0.1-1.1-0.5-1.2    c-0.9-0.1-1.6,0.5-1.6,1.2c-0.1,1.7,0,3.5,0,6c-1.8-2.1-3.1-4.2-2.2-5.8c1.3-2.4-0.5-3.4-1.1-5c-0.5-1.6-2.1,0.4-2.8-0.7    c-0.2,0.8-0.8,1.4,0,2.5c1.5,2.2,1.3,4.6-0.9,7.5c0.3-3.6-0.2-6.4-1.9-8.8c-1.3-1.9,0.4-3.4,1.2-3.8c3-1.4,3.6-4.4,4.2-6.6    c0.8-3.2,2.6-5.7,3.7-8.8c-5,2.7-4.6,10.4-11,11.8c1.1-6,6.3-10.2,7.3-16.4c-5.9,7.8-8.3,18.4-18.7,23.3c1.1-4.9,8.1-5.9,6.8-11.8    c-1.8,4.1-5.3,6.6-8.3,9.4c-1.5,1.5-2.2-0.9-3.1-1.7c-0.9-0.7-1.6-1-1.9,0.1c-0.7,2.5-3.2,3.1-5.9,1.3c-0.5-0.4-0.9-0.9-1.7-1.5    c-0.9,4.4-4.4,5.1-7.8,5.8c-0.5-1.3,1.1-1.5,1.1-2.5c-1.1-0.8-2,0.3-2.3,0.6c-3.3,3.6-8.8,3.6-11.9,7.5c-1.3,1.7-3-0.5-4.6,0.3    c1,1.8,3.2,1.2,4.3,2.4c-1.3,2.3-3.9,2.9-6.1,3.8c-1.8,0.8-2.5-0.7-1.9-2.3c0.6-1.4,1.7-2.7,2.7-4.1c-4.5-0.4-8.5,1.2-10.2,4.9    c-2.6,5.5-6.7,5.9-11.9,5.6c0.9-2.9,2.4-5.4,5-5.8c1.6-0.2,2.2,2.3,2.1,4.2c2.1-1.2,3.2-2.7,0.7-3.9c-3-1.5-2-3.2-0.9-5.2    c0.5-0.8,1.4-1.6,0.5-2.8c-1.1,0.6-2.1,1.6-2.8,2.8c-0.5,1-0.8,1.9-2.5,1.1c-1.8-0.9-2.9,1.2-3.3,2.3c-1.3,3.5-3.8,3.8-7.3,3.4    c3.8-4.4,8.1-7.8,13.4-9.9c4.2-1.7,8.2-3.7,11.1-7.2c4.4-5.3,10.4-8.5,16.6-11c4.7-1.9,8.3-5.2,12.5-7.7c3.1-1.9,6.2-3.5,9.5-4.7    c2.8-1,5.4-2.3,7-5.3c0.9,1.9-0.5,4.1,2.1,5.6c1.5,0.9-1.4,2.1-2.5,3.2c1.9,0.2,3.6,0.5,5.3,0.6c0.9,0.1,2-1.9,2.7-0.2    c0.4,0.9-1.3,1.5-2,2.2c-1.5,1.7-3.8,2.2-5,4.4c4-1.9,8.1-3.9,12.1-5.8l0.2-0.1l-0.1,0.2c-1.5-1-5.8-0.3-2.7-4.6    c-0.6,0.3-1.3,0.7-2.2,1.2c0.6-2.5,1.1-4.8,1.7-7.3c-3.9,1.5-0.7,7.7-6.4,8.3c2.2-2.5,3.1-4.8,2.3-7.5c-2.2,0.3-0.5,2.8-2,3.6    c-1.7-3.8-0.6-8.9,2.7-12.4c1.9-2,4-3.8,5.4-6.2c0.4,0,0.7-0.2,1.1-0.1c-0.2-0.1-0.4,0-0.6,0.1c-0.5,0.2,0.1,0-0.3,0.2    c-3.5,1.3-6.6,3.2-9.1,5.9c-1,1.1-2.1,1.5-3.7,1.6c2.5-2.3,2.8-5.8,5.4-8c0.4-0.3,2.9-1.3-0.2-1.8c-1.3-0.2,0.2-1.2,0.5-1.4    c2.1-1.3,4.1-2.7,6.6-3.1c1.2-0.2,2.3-0.6,3.1-1.8c-1.7-0.9-3,1.2-4.5,0c2-1.8,4.1-3.5,6.8-4.3c2.5-0.7,5-1.4,7.4-2.4    c-3.5-0.1-7.4,1.3-9.9-2.4c6.6,4.2,13-2.9,19.7-0.3c-1.4,1.2-3.6-0.4-4.2,1.8c2.6,1.2,4.7-2,7.9-0.8c-3.1,2.1-6.4,2.3-9.3,3.7    c2.8,0.5,14.5-3.3,16.3-5.3c-3-2.8-7.5-1.6-10.7-4.1c-6.5-5.2-15.7-5.1-22.1-10.7c-0.7-0.6-1.4-0.9-1.6-2.6c2.4,2.1,4.7,1.8,7,1.8    c0.8,0,1.8-0.4,2.5,0.6c3.4,4.8,8.8,3.7,13.5,4.9c3.3,0.9,7.6,0,9.8,4.2c0.4,0.7,2.8,1.7,4.6,0.8c-1.2-0.6-2.3-1.1-3.9-1.8    c2.8-0.2,4.9-0.9,6.9-2c1.4-0.8,2.2-0.5,3.3,0.4c1.3,1.2,0,1.6-0.4,2.3c-0.4,0.7-0.6,1.6,0.3,2c0.8,0.3,1.3-0.5,1.7-1.2    c2-2.9,0.4-5.6-1.5-7.3c-5.1-5-9.2-11.2-16.2-14c-1.1-0.5-1.3-0.7-0.4-1.9c2.1-2.8,3.1-6.3,0.7-9.2c-3-3.7-4.2-5.4-10.1-4.3    c0.2,1.9,1.8,2.7,2.9,3.7c2.7,2.5,1.9,5.2,0.7,7.9c-0.4,0.9-1.2,1.4-2.2,1.1c-1.2-0.5,0.2-1.1,0.1-1.7c-0.1-1.4-1.1-2-2-2.6    c-0.8-0.5-1.6,0-2.4,0.6c-1.2,1-2,3.3-3.7,2.9c-3.7-0.9-7.8-0.1-11.1-2.5c-1.8-1.3-2.8-2.5-0.4-4.5c2.5-2.1,5.1-4.2,6.4-7.4    c0.6-1.6,2-0.8,2.9-0.8c8.5-0.5,16.9,1.2,25.3,1.4c7.7,0.2,15.2-2,22.5-4.2c8.1-2.5,16.3-2.2,24.5-0.9c1.9,0.3,3.8,1.2,4.8,3.2    c-4.8,2.6-10.1,4-15.7,4.2c5.8,2.8,10.8,2,16.3-0.8c3.1-1.6,7.6-1.2,11.4,0.3c4.3,1.8,8.6,3.5,12.9,5.3c0.8,0.3,1.8,0.7,2.1,1.3    c0.4,1-0.7,1.6-1.4,2c-6.7,5-13.9,8.9-22.4,9.5c-3.3,0.2-6.5-1.5-9.1-3.5c-1.5-1.1-2.7-2.4-4.7-1.4c0.8,2.9,2.2,5.6,5.1,7.4    c-1.8,1.6-2.1-0.6-3.1-0.6c-1.4,2.4,1.7,2.1,2.1,3.7c0.6,2.3-0.3,2.7-2.5,2.2c1.5,2.7,2.8,5,4.4,7.8c-2.4-0.2-2.2-2.5-3.9-2.8    c-0.8,0.4-0.5,1.3-0.3,1.9c1.1,4.3-2.8,5.7-4.8,8.1c-1.2,1.4-4.9-0.5-5.8-2.7c-1.4,0.5-1.2,2.1-2,2.8c-0.7,0.6-1.4,3.6-3.1,0.6    c-0.4-0.7-1.1-0.1-1.6,0.3c-1.1,0.9-2.2,1.6-3.6,1.7c-0.2-0.4-0.4-1.3-0.8-1.1c-5,1.5-9.9-1.3-15-0.4c-0.4,0.1-0.9,0.2-0.7,0.1    c3.2,0,6.9-1,10.4,1.4c1.9,1.4,4.1-0.5,6.2-0.1c0,0.6-0.2,1.5,0.6,1.5c3.2,0,5.6,2.9,9,2.3c-0.2-1.3-5.7-1.7-0.7-3.9    c1.3-0.6,1.4-0.6,1.9,0.8c0.8,2,3.3,1.6,4.6,0c1.4-1.8,2.9-3.5,5.1-4.4c0.6-0.2,1.3-0.5,1.5-1c2-4.2,3.3-1.9,4.8,0.4    c0.9,1.4,1.8,2.9,4.4,3.9c-1.6-2.5-2.7-4.4-4-6.2c-1.2-1.6-2-9-1-10.6c0.6-1,1.3-0.9,2,0c1.4,1.6,3.1,2.9,4.3,4.7    c0.2-4.8-3.7-7.6-6.2-11.8c4.4,1.4,6.4,4.3,7.2,7.7c0.7,2.7,1.3,5.6,0.9,8.5c-0.3,2.7,0,5.3,1.6,7.5c1.3,1.8,1,3.8-0.3,4.8    c-3.6,2.7-1.4,5,0,7.6c0.5,0.9,1.5,1.8,0.7,3.3c-0.8-0.7-1-2.3-2.7-1.9c0.8,3.5,2.8,6.6,4.2,9.8c0.4,0.9,0.6,2,0.1,2.5    c-0.7,0.6-0.4-1.1-1.4-1c-0.9,0.2-1.6-0.1-1.2,1.4c1.3,6,1.3,12,1.3,18.4c-1.5-1.3-1.3-3.4-3.1-3.8c-0.1-0.1-0.3-0.2-0.4-0.3    c-1.2-1.4-1.9-3.2-3.6-4.3c-0.2,0-0.3,0-0.5-0.1c-2.8,0.6-1.5,2-0.9,3.6c0.7,1.9,2.4,3.6,1.4,6l0.1-0.1c-0.4,0-0.7,0-1.1,0    c0.9,2.3,0.8,5.4,2.9,6.6c3.5,2,3.6,5.6,5.1,8.5c0.7,1.4,3,2.9,1.3,5.3c-1.9-4-4-7.8-7.5-10.8c-0.4,6.1,7,8.4,6.3,15.1    c-1-2.6-3.6-3.2-3.7-5.7c0-0.8-0.6-0.8-1.1-0.6c-0.2,0.1-0.5,0.7-0.4,0.9c0.9,1.8,0.5,3.7,0.9,5.5c0.1,0.4,1.1,1.1,1.3,1    c3.7-1.6,3.1,1.6,2.9,3.1c-0.2,2.3,2.2,3.7,1.4,6.3c-1.8-2-3-4.4-5.5-5.7c-1,2.7,1.2,4.3,1.2,6.3c1.6-0.2,2.5,0.2,3.2,2    c1.5,4,2.4,8,2.1,12.3c-1.1-0.3-0.8-1.7-2.3-1.7c0.7,2.9,1.3,5.7,2,8.7c-1.5,0.7-1.8-1.7-3.1-1.3c-2,2.3,6.1,3.2,0.7,6.3    c-0.5-3.5-3.4-5.8-3.9-9.1c-0.4,5.2,2.8,9.2,4.9,13.5c0.6,1.2,1.1,2.2,1,3.5c-0.5,4.3,1.1,8.3,2.2,12.4c0.9,3.4,2.9,5.2,6,6.5    c9.4,3.9,18.1,2.4,25.7-3.9c5.2-4.3,3.6-12.3-2.3-14.9c-2.2-1-5.9,1.9-6.3,4.9c-0.4,2.9,1,5.3,2.6,7.6c0.6,0.9,2.4,2.3,0.5,3.4    c-2.1,1.3-2.1-1.1-2.7-2c-1-1.8-1.7-3.8-2.6-5.7c-1.9-3.7,0-6.5,2.3-9c1.9-2.1,4.4-2.7,7.1-1.4c3.4,1.5,5.4,4,5.9,7.7    c1.2,9.5-5.5,14.8-13.7,16.8c-8.3,2-16.5,0.7-23.1-5.6c-0.4-0.4-1-0.7-1.5-1.1c-4.4-3-5.2-8.8-1.9-12.9c0.7-0.8,1.8-1.9,1.1-2.5    c-3.1-2.8-1.9-7.4-5.2-10.3c0,4.1,2.1,7.8,1.3,12.4c-0.7-1.8-1.4-2.8-3-3.2c3.1,2.7,0.6,6.5,2.3,9.5c0.3,0.6-0.1,0.7-0.8,0.5    c-2.2-0.7-1.3,0.4-0.9,1.5c1.1,2.7,1.1,5.6,0.9,8.9c-2.8-2.7-4.4-2.6-4.7,1.7c-0.1,1.7-1.8,3.3-2.5,5.1c1.9-1.1,2.8-3,3.9-4.8    c0.4,3.1-1.5,5.2-3.4,7.2C206.1,249.1,204.6,249.3,204.7,250.8c-3.9,0.3-7.4,2.2-10.9,3.9C197.3,252.9,200.8,251.4,204.7,250.8z"
		pathData2 := "M138.9,217.2c-2.6,3.5-3.9,7.9-3.1,11.7c1.7,8.7-1.2,15.8-5.7,22.6c-0.4,0.7-1.7,1.3-2.5,1.2    c-1-0.2-0.4-1.7-1.1-2.6c-0.6-0.8,0.5-1.4,1.3-1.9c4.8-2.4,5.5-4.2,3.1-9c-1.7-3.3-0.2-6.5-0.3-9.8c-1.1,3.7-3.4,7.2-0.5,11.2    c2.1,3,1.2,4.4-2.6,6c-1.1-1.4,0.9-1.7,0.9-2.8c-0.2-0.1-0.6-0.3-0.7-0.2c-1.6,1.9-3.3,1.6-5.3,0.4c-1.5-1-3.3-1.6-5-2.2    c-5.2-1.9-5.5-3.7-1.9-8c-2.3-1.1-3.7-3.3-6.2-4c-1.7-0.5-1.8-2.4-1.8-3.4c-0.2-3.2-2.3-4.5-4.7-5.7c-5-2.6-5.2-3.9-1.2-7.4    c-0.7-0.3-1.3-0.2-2.1,0c-2.2,0.6-4.3,1.1-3.2-2.7c0.6-1.9,0-4.1,1.5-5.9c1.5,0.4,2.7,1.4,4,2.3c-0.9-1.6-2.5-1.8-4.1-2.2    c1.3-1.6-2.7-1-1.2-3.2c1.8-2.7,3.8-1.1,5.2-0.4c3,1.3,4.8,3.9,6.3,7c2.6,5.3,6,10.2,10.9,13.8c1.1,0.8,2.4,1.2,3.7,1.1    c-1,4.2,0.5,8.4-0.5,12.6c1.1-4.2,2.2-8.4,3.5-13.5c3.1,7.7-1.4,14-0.8,20.9c2.8-5.3,3.1-10.9,3.4-16.1c0.3-4.5,2.4-7.3,4.9-10.4    c0.1-0.1,0.4-0.1,0.6-0.2c0.1,4.1-3.7,7.3-2.7,11.8c2.8-6.2,4-13.1,9.5-17.7c0.6-0.5,0.9-1.6,1.9-1.1c1.5,0.7,0.1,1.4-0.1,2.1    c-2.9,10.8-5,21.6-0.3,32.4c0.4,1-0.4,2.3,1.5,3.3c-3.5-13.4-3.4-26.2,2.7-38.5c1.4,3.9-3.5,7.1-1.3,11.1c1.8-2.2,2.2-5,2.7-7.7    c0.5-3.1,2.5-4.7,5.4-4.2c1.7,0.3,1.6,1.7,1.5,3c-1.3,9.3-1.9,18.6-1.4,27.9c0.1,2.9,0.6,5.9,2.1,8.6c1.6-1.9,1.7-2.1,1-2.9    c-1.8-1.9-1.7-4.1-1.7-6.3c0-7.7-0.2-15.3,0.7-22.9c0.3-2.1,1.3-3.5,2.4-5.5c1.3,2.1,0.2,3.9,1,5.6c1.1-0.8-0.2-2.7,1.9-3.3    c0.1,0.9,0.2,1.8,0.2,2.6c-0.2,8-0.5,15.9-0.8,23.9c-0.1,2.5,0.8,4.4,2.8,6.8c-2.2-10.9-1.1-21.1-0.5-31.3    c0.1-1.5,0.7-1.6,1.9-1.6c2,0.1,2.7,2.1,2.6,3.1c-0.5,3.7,0.5,7.5-0.5,11.2c-0.4,1.7-0.2,1.8,1.9,4.2c0.3,0.4,0.4,0.8-0.1,1.3    c-3.2,3.6-5.8,7.4-3.6,12.5c0.1,0.3,0.1,0.7,0,1c-0.3,0.6-0.7,1.2-1.1,2c-2.4-2.4-4.1-5-4.2-8.9c-1.7,4.7-1.7,4.7,0.8,9    c-0.5,0.9-1.3,0.4-2,0.5c-3,0.4-4.7,0-5.7-3.8c-2.2-8.6-2.6-17.2-1.6-25.9c0.3-2.5,1.1-4.9,1.6-7.3c0.2-0.8,0.7-1.8-0.7-2    c-1.6-0.4-0.8,0.9-1,1.5c-2.5,7.3-1.9,15-1.7,22.4c0.2,5.4-0.4,11.3,3.2,16.2c0.4,0.6-0.4,2-0.9,4.2c-0.2-4.6-2.9-7-3.7-10.6    c-0.8-3.6-2.2-7-2-10.9c0.2-3.7,0-7.4,0-11.3c-2.9,2.3-2.1,5.9-1.9,8.1c0.3,4.6,0.1,9.4,2.2,13.8c1.4,3,0.7,3.7-2.9,4.9    c-1.3,0.5-2.9,1.7-2.8,3.7c0.1,1.4,0,2.8,0,4.4c-2.2-1.3-3.7-2.5-1.5-4.7c2.8-2.8-0.1-5.9,0.9-8.9c-1.7,0.4-2.1,1.6-1.6,2.5    c1.8,3.3,0,5.5-2.2,7.5c-2.4,2.2-4.2,4.6-4,8.5c-2-2.2-2.6-4.1-1.1-6.4c3-4.8,5.5-9.9,7.1-15.3c0.8-2.6,0.7-5.9-0.1-8.5    c-0.8-2.5-0.6-4.8,0.3-6.9c1-2.3,0.2-4.5,0.8-6.7c0-0.2,0.1-0.4,0.3-0.5C138.6,216.6,139.1,217,138.9,217.2z"
		pathData3 := "M112.5,121.4c-4.5,0-8.9-1.1-12.8-3c-3.7-1.8-7.5-3.9-9.6-7.8c-0.5-0.9-2.4-1.6-0.1-2.7    c0.5-0.2-0.4-1-0.8-1.3c-2.4-1.6-3.6-4.1-4.8-6.9c1.9,0.3,2.5,1.6,3.3,2.7c0.1-2.5-3.6-7-7.5-7.7c-3-0.6-3.4-2.2-3.6-4.5    c-0.1-0.6,0-1.4-0.3-1.8c-0.5-0.7-0.5-2.3-2-1.6c-1.3,0.6,0.1,1.4,0.1,2.1c0.1,1.2,1,2.3,0.4,3.8c-2.3-1.7-4.2-3.8-6.1-6    c4.3-1.4,11-1,12.9,0.6c1.1-0.6,1-2.3,2.5-2.6c-3.1-1.3-6-2.5-7.1-6.2c-0.8-2.8-5.7-4.7-1-8c-2.4-2-4.1,2.9-6.3-0.1    c-1.7-2.3-0.2-4.7-1.3-7.7c3.9,2.5,5.8,6.5,10,6.9c-0.1-0.9-0.3-1.9-0.3-2.8c-0.1-1.1,1.1-1,1.6-1.2c1-0.3,0.5,0.8,0.6,1.2    c2,7.2,6.3,12.4,12.4,16.9c6.6,4.9,10.6,12.4,15,19.4c2.3,3.7,4.8,7.2,8.5,9.6c0.7,0.5,0.6,0.6-1.4,2.2c1.9,0,3.4,0,5.7,0    c-2.7,2.6-5.1,3.3-8.1,2c-4-1.7-8.2-3-12.5-4c3.3,3,8.2,2.7,11.6,5.6c-5.7,0-10.2-3.2-15-5.4c-1-0.5,0.4-3.4-2.2-3.6    c-0.7-0.1-0.9-1.8-2-0.6c-0.8,0.9-0.6,1.7,0.1,2.7c2.5,3,6.2,4.1,9.4,5.9C105.3,119.2,109.1,119.6,112.5,121.4z"
		pathData4 := "M246.2,280c4-4.1,5.4-8,3-12.9c-1.2-2.4-0.1-5.5,1.9-8.3c-3.3,1.7-3.3,1.7-4,7.6c-2.7-3.3-2.3-6.3,1.1-8.9    c3.2-2.5,6.7-5,6.1-10c-0.3-2.3-2.5-3.6-2.9-6.4c4.1,3.3,7.8,6.1,7.4,11.6c-0.3,3.6,1.9,6.3,2.7,9.5c0.3,1.1,1,2.4-0.3,2.9    c-1.1,0.4-1.2-1.3-1.9-1.8c-1.2-1-1.1-3.3-3.3-3.3c-0.4,1.8,0.8,3.3,1.6,4.3c2.5,2.9,1.4,5.3-0.3,7.8    C254.8,275.7,251.9,279.1,246.2,280z"
		pathData5 := "M112.7,134.9c-0.9,3.9-3.7,4.5-6.4,5.3c-1.2,0.4-1.7-0.4-2.1-1.1c-1.7-3.1-3.4-6.1-6.2-8.3    c-0.7-0.6-1-1.3-0.6-2c0.7-1.4,2.4-0.1,3.4-0.9c-0.9-1-2.1-0.1-3-0.6c2.3-4.8,6.7-1.3,10.9-2.7c-2.9,0-4.6-1-6.3-1.8    c-0.3-0.2-0.7-0.7-0.9-0.6c-6.1,2.9-5.4-4.4-8.8-5.6c5.9,3.5,12.4,5.5,19.2,7c-1.1,1.5-2.9,1.3-3.7,2.8c3.9,0.8-0.4,1.6,0,2.3    c1.5,1.2,2.4-0.5,4-1.4c-2.1,4.7-7.1,4.2-9.9,7.2c4,0.6,7.3,0.3,10.1-2.5c-0.8,3.4-4.7,3.1-5.8,6.4    C108.6,137.1,110.4,136.1,112.7,134.9z"
		pathData6 := "M191.3,51c3.5-5,7.1-8.6,11.2-11.7c2.5-1.9,5.2-1.9,7.8-1.4c2.1,0.4,2.7,3,1.2,4.6    C206.3,48.3,199.3,50,191.3,51z"
		pathData7 := "M156.5,87.4c1.8,0,3.6,0,5.5,0c-0.8,1.6-2.4,0-3.2,1.4c1.4,0.1,2.7,0.2,4,0.3c-0.2,1.2-1.5,0.1-1.7,1.2    c0.4,0.9,2.5-0.6,2.1,0.9c-0.2,1-1.2,2.4-2.9,2.4c-2.1,0-3.8-1-5.7-1.5c-3.5-0.8-6.8-2-10.3-2.9c-3.4-0.9-3.3-3-3.5-5.7    c6.1,2.7,12.3,1.2,18,2.8C158.6,87.4,157,86,156.5,87.4z"
		pathData8 := "M129.3,85.2c-2.9-0.3-6.5-3.1-10.8-2.4c-0.7,0.1-1.5-0.8-1.2-1.5c0.4-1,1.1-2,2.4-2.1    c2.9-0.2,5.7,0.2,8.4,1.2c1.1,0.4,2.3,1,3.5,1c1.6,0,2.5,0.1,2.5,2.2c-0.1,2.2-1.4,1.6-2.6,1.7C131.2,85.2,130.8,85.2,129.3,85.2z    "
		pathData9 := "M122.6,59.9c-0.7,0.4-1.4,0.8-2.2,1.3c-0.5-2.1,0.5-3.7,1.2-5.4c0.8-2.1,2.1-3.6,4.5-2.6    c2.1,0.8,3.1,4.5,1.6,6.3c-1.5,1.9-3.3,3.5-5,5.2c-0.5,0.5-0.9,1.9-2,0.6c-0.7-0.8-0.6-1.6,0-2.4c0.7-1,2.2-1.4,1.8-3.1l0.2-0.1    C122.7,59.7,122.6,59.9,122.6,59.9z"
		pathData10 := "M251.5,233.8c-1.2-2.3-2.7-4.1-3.5-6.5c-1-2.7,0.3-5,1.8-5.5c2.1-0.7,4,1.8,4.4,3.7    C254.8,228.4,254.5,231.5,251.5,233.8z"
		pathData11 := "M211.7,48.3c0.6-0.5,1.2-0.9,2.1-1.5c-0.4,3.4-3.1,3-5,3.8c1.2,0,2.4,0,3.6,0c0.7,0,1.2-1.3,1.9-0.4    c0.5,0.7,0.5,1.8,0,2.5c-0.6,0.7-1.4,2.2-2.3,1.6c-2.7-1.8-6-1.7-8.9-3.4c3.1-2.6,7.1-3.4,9.6-6.4C213.7,46.3,211,46.5,211.7,48.3    z"
		pathData12 := "M154.4,83.4c-2.9-0.9-5.8-0.6-8.6-0.5c-1.6,0-2.9-0.5-4.2-1.3c-0.9-0.6-0.7-1.6-0.6-2.4    c0.1-1.2,1-0.8,1.7-0.9C147.5,77.8,151.3,79.7,154.4,83.4z"
		pathData13 := "M145.5,62.3c4.7,1.3,6.9,4.3,6.2,8c-0.7,3.8-4.9,6.6-9.9,5.9c3-1.1,6-1.8,7.2-5.5    C150.3,66.8,147.6,64.9,145.5,62.3z"
		pathData14 := "M142.3,70.9c-3.4,6.5-5.4,12.3-1.9,18.8C134.3,87.1,135.2,76.8,142.3,70.9z"
		pathData15 := "M154.9,192.8c3.4-0.1,3.8,2.7,4.9,4.5c0.3,0.5-1,1.6-1.9,1.9c-0.7,0.3-4.6-4.4-4.5-5.5    C153.6,192.8,154.1,192.1,154.9,192.8z"
		pathData16 := "M88.4,202.5c2,2,8.7,0,5.5,6.5C91.3,207.6,88.9,206,88.4,202.5z"
		pathData17 := "M215.2,144.9c0,1.5,1.6,2.9,0.2,5c-1.5-3.5-4.5-5.8-4.7-9.6c0,0-0.1,0.1-0.1,0.1c3,0.1,3.2,3,4.7,4.6    L215.2,144.9z"
		pathData18 := "M74.5,56.5c1.4,2.9,3.6,4.1,5.2,6.3c-1.5-0.4-1.9,2.1-3.2,1C74.4,62.1,74.1,59.7,74.5,56.5z"
		pathData19 := "M211.2,149.8c0.7,0,2-0.3,2.1-0.1c1.6,3.2,5.5,5.2,4.8,9.9c-1.9-3.9-4.2-7.1-6.8-10    C211.3,149.7,211.2,149.8,211.2,149.8z"
		pathData20 := "M91.2,195.8c1.4,2.1,3.7,2.3,4.1,4.3c0.1,0.4-1,1.8-2.3,1.4C90.4,200.6,91.3,198.4,91.2,195.8z"
		pathData21 := "M92.4,215.7c2.6-0.6,4.2-0.4,4.8,1.7c0.1,0.5-0.1,1.7-0.2,1.7C94.8,219,94.3,216.9,92.4,215.7z"
		pathData22 := "M171.3,204.8c0.6,1.4,0.1,2.2-0.8,2.8c-0.1,0.1-0.8-0.4-0.8-0.6C169.5,205.8,170.5,205.5,171.3,204.8z"
		pathData23 := "M234.8,68.3c-6.1-0.2-11.4,0.5-17.1-0.7c2.3,2.4,5.7,1.5,8.3,3.8c-4.8-0.3-8.9,0.3-12.8,2.2c-1.1,0.6-2.5,0.1-3.4-1.1    c-0.9-1.2-1.9-2.3-3.8-2.3c3.4-0.9,5-3.4,6.7-5.9c2.5-3.6,6.5-2.7,9.4-1.9C226.2,63.6,230.6,64.6,234.8,68.3z"
		pathData24 := "M172.1,67.1c-1.3-1-3.2,0.5-4.7-1.5c7.8-0.9,15.1,0.2,22.5,1.4c-1.4,2-3,1.3-4.8,0.9c-1.4-0.3-4.6-2-2.9,2.2    c0.6,1.5-0.6,2.6-1.8,3.4c-1.8,1.1-4.1,1.8-5.6,0.5c-2.6-2-5.7-2.8-8.3-4.5C167.6,66.8,171,69.8,172.1,67.1z"
		pathData25 := "M181.6,104.2c-2.7,1.8-4.8,2.8-8.1-0.1c-2.1-1.8-6.3-0.9-9.5-0.2c-1.1,0.2-2.1,0.6-3.6-0.4c4.1-1.5,8-2.3,11.7-1.2    c3.8,1.1,7.1,0.2,10.6-1c1.1,1.6-1.2,1.9-1.2,3.1C181.5,104.4,181.6,104.2,181.6,104.2z"
		pathData26 := "M138,160.6c1-3,0.8-6.2,4.2-8.5c2.3-1.6,3.3-5,4.9-7.7C148.3,147.6,143.7,155.9,138,160.6z"
		pathData27 := "M154.1,139c-1.5,2-1.9,4.7-3.5,6.7c-1.7,2.2,0.7,4.3,0.4,6.8c1.3-0.8,0-2.5,1.3-3l-0.1-0.1c0.4,3.9,0.3,7.7-0.7,11.5    c-0.7-1.7-2.2-3.3-0.7-5.3c-1.8-0.5-1.3-2-1.6-3.3c-1.5-5.2,1.3-9.1,4-13.1c0.1-0.2,0.6-0.1,1-0.1L154.1,139z"
		pathData28 := "M177.2,179.6c1.9,4.5-2.6,6-4.8,8.8c0.8-6.8,1.7-13.7,2.5-20.5c0.3,0,0.7,0.1,1,0.1c-0.6,5.3-1.2,10.5-1.9,16    C176.6,183.4,174.8,180,177.2,179.6L177.2,179.6z"
		pathData29 := "M105.5,164.8c1.1-5,6.8-7,7.6-12.3c0.8,2.1-0.4,3.6-1.3,5.7c3.4-1.5,4.9-4.5,7.3-6.1c-0.5,2.5-3.7,6.8-6.2,7.3    C109.5,160,107.3,161.9,105.5,164.8z"
		pathData30 := "M208,232.1c-1.3-1.5,3.7-2.9-0.6-3.8c-0.4-0.1-0.4-0.4-0.1-0.8c3.1-3.4,0.8-8.2,3.3-11.8    C209.1,221.1,213.8,227.3,208,232.1z"
		pathData31 := "M165.2,177.7c0.9,6.1,1.8,11.8,2.7,17.9C165.5,193,163.7,181,165.2,177.7z"
		pathData32 := "M216.2,260.2c3.1,1.8,6.2,3.6,6,7.8c0,1.1-0.2,2.1-1.3,2.4c-1,0.3-1.5-0.8-2.1-1.4c-0.4-0.4-0.7-1-0.9-1.6    c-0.3-0.9-2.7-1.3-1.3-2.4c1.3-0.9,1.8,1,2.5,2c0.5,0.6,0.3,2.6,1.8,1.7c1-0.6-0.1-1.7-0.3-2.6    C220.1,263.4,217.5,262.3,216.2,260.2z"
		pathData33 := "M204.7,250.8c-2.1,1.7-4.8,1.8-7.1,2.9c-2.2,1-4.3,2.1-6.8,3.4C192.2,253.4,200.5,250,204.7,250.8    C204.8,250.9,204.7,250.8,204.7,250.8z"
		pathData34 := "M156.9,152.1c0.3,5.1-0.6,9.2-1.6,13.2C152.9,160.9,155.3,157,156.9,152.1z"
		pathData35 := "M194.9,57.4c4.3-1.1,8.4-2.7,13.8-1.9C203.6,56.7,199.5,58.9,194.9,57.4z"
		pathData36 := "M161.5,179.4c0.8,0,0.9,0.5,0.9,1c0,3,0.9,5.8,1.5,8.7c0.1,0.6,0.2,1.4-0.4,1.7c-0.8,0.3-1.6-0.4-1.7-1.1    c-0.4-3.2-0.6-6.3-0.9-9.5C160.9,179.9,161.3,179.6,161.5,179.4z"
		pathData37 := "M147.4,154.3c-4.5,2.2-5.2,7.8-9.5,10.2C140.9,160.8,142.1,155.7,147.4,154.3z"
		pathData38 := "M172.1,166.6c-0.1,5.4-2,10.6-1.5,16.1C170.3,177.2,169.1,171.7,172.1,166.6z"
		pathData39 := "M190.1,204.7c-1.2-2.5-0.3-5.4-1.6-7.8c-0.1-0.3,0.1-1,0.4-1.2c0.4-0.3,0.7,0.3,1.1,0.6    C193.6,199.1,189.4,201.8,190.1,204.7z"
		pathData40 := "M170.6,72.1c0,3.2,0.8,5.6,4,7.3c-3.8,0.6-4.6-2-5.8-3.7C167.7,74.1,169.3,73.3,170.6,72.1z"
		pathData41 := "M156.6,136.1c-1.7-2.9,2.8-2.6,2.4-4.9c0.5,0,1.2,0.1,1,0.7c-1.1,2.5-1,5.6-3.4,7.8C155.8,138.1,158.1,137.1,156.6,136.1    L156.6,136.1z"
		pathData42 := "M158.7,147.1c-0.2-3.4-0.7-5.9,0.9-8.2c0.3-0.5,0.9-1.1,1.5-0.7c0.7,0.5,0.3,1-0.1,1.6    C159.7,141.6,160.4,144.2,158.7,147.1z"
		pathData43 := "M159.1,159.9c0,2.8,1.1,5.7-1.9,8.8C157.1,165,157.7,162.4,159.1,159.9z"
		pathData44 := "M211.9,110.9c-0.9-2.5-3.8-4.1-3.3-7.7C211.2,105.5,212.4,107.8,211.9,110.9z"
		pathData45 := "M126.2,140.9c1.1-2.6,1.4-5.5,5.4-5.1C129.7,137.5,129.2,140.2,126.2,140.9z"
		pathData46 := "M97.2,163.2c1.2-1.5,2.3-2.8,4.7-2.4c1,0.2,2.3-1.7,3.4-2.6c-0.7,3.1-2.2,4-8.2,4.9L97.2,163.2z"
		pathData47 := "M183.3,186.6c0.3,3,1.4,6.2-1.7,9.1C182.2,192.3,182.8,189.5,183.3,186.6z"
		pathData48 := "M160.5,171.5c-0.8-3.6,0.4-5.8,1.3-8.1C161.7,165.9,162,168.5,160.5,171.5z"
		pathData49 := "M167.2,177.9c-1.6-2.7-0.5-4.5-0.2-6.3c0.1-0.7,0.7-0.8,1.1-0.5c0.3,0.2,0.5,0.9,0.4,1.3    C168.2,174,167.8,175.5,167.2,177.9z"
		pathData50 := "M195.7,245.4c3.8,1.1,5-2,6.9-4.1C202.4,245.8,201.4,246.4,195.7,245.4z"
		pathData51 := "M190,112c0-1,0-1.6,0.9-1.6c1.9,0.1,2.7,1.6,3.6,2.9c0.3,0.4-0.3,1.1-0.8,0.8C192.4,113.5,191.1,112.7,190,112z"
		pathData52 := "M163.4,175.8c-0.2-3.3-0.4-5.9,0.1-8.8C164.7,168.8,164.7,168.8,163.4,175.8z"
		pathData53 := "M214.3,257.8c-1.3,2.1-3.1,2.4-5.7,1.6C210.5,257.5,212.8,259.1,214.3,257.8z"
		pathData54 := "M215.3,145c-2.1-1-2.2-4-4.7-4.6c0.2-0.7,0.3-2.5,1.2-1.2C213.2,140.9,215.3,142.4,215.3,145z"
		pathData55 := "M199.1,85.9c-2-1.8-4.1-2.5-3.6-5.3C197.2,81.8,198.4,83,199.1,85.9z"
		pathData56 := "M177.3,179.7c-0.4-2.4-0.9-4.8,0.7-7.6C179.1,175.2,177.7,177.3,177.3,179.7C177.2,179.6,177.3,179.7,177.3,179.7z"
		pathData57 := "M142.8,165.8c0.9,2.3,0,4.3-0.7,6.4C140,169.7,143.2,168,142.8,165.8z"
		pathData58 := "M165.9,64.4c2.2-1.6,3.9-3,6.5-2.8C170.4,63.1,168.7,64.2,165.9,64.4z"
		pathData59 := "M172.2,101.1c-0.4-1-0.7-1.3-0.6-1.4c0.9-1.1,0.4-2.8,1.7-3.6c0.1-0.1,0.9,0.3,0.9,0.5C174.2,98.2,173.2,99.3,172.2,101.1    z"
		pathData60 := "M119.2,156.6c-1.3,1.6-2.4,3.6-4.2,4.3C116.2,159.6,117.7,158.1,119.2,156.6z"
		pathData61 := "M189.5,92.2c1.6-0.5,2,0.4,2.3,1.4c0.3,1.2-0.7,0.6-1.1,0.7c-1.1,0.3-1.4-0.5-1.8-1.3C188.3,91.7,189.4,92.3,189.5,92.2z"
		pathData62 := "M119.4,154.9c1-1.2,1.9-2.5,3-3.8C123.1,153.8,120.9,154.1,119.4,154.9z"
		pathData63 := "M113,164.5c0.1-1.3-2.3-1.4-1.3-2.9c0.3-0.4,0.8-0.9,1.2-0.5C113.9,162.1,113.3,163.4,113,164.5z"
		pathData64 := "M175.2,101c1.3-1.9,1.7-3.4,3-4.3C178.9,99,177.5,99.7,175.2,101z"
		pathData65 := "M154.2,145.9c-1.1,0.9,1.5,3.9-2,3.4c0,0,0.1,0.1,0.1,0.1C152.6,148.1,152.1,146.3,154.2,145.9L154.2,145.9z"
		pathData66 := "M184.4,200.4c0.1,0.4,0.2,0.7,0.2,1.1c-0.1,1.1,0.6,2.4-0.5,3.1c-0.7,0.5-1.3-0.2-0.8-1.2c0.3-0.5,0.4-1.1,0.5-1.7    C183.9,201.1,183.8,200.6,184.4,200.4z"
		pathData67 := "M126.3,121.8c0.8-1.3,2.2-1.7,3.8-2.4c-0.9,2.2-2,2.9-3.9,2.3C126.2,121.7,126.3,121.8,126.3,121.8z"
		pathData68 := "M129.7,117.6c1.4-1.6,1.4-1.6,4.3-2.4c-1.1,1.7-2.8,2-4.4,2.3L129.7,117.6z"
		pathData69 := "M169.4,126.4c0.9-1.6,2.3-1.5,3.4-2.1C172.2,125.8,171.4,126.9,169.4,126.4z"
		pathData70 := "M138.8,167c0.6-1.1,0.6-2.3,1.8-2.8c0.6,0.6,0.5,1.3,0.1,1.9C140.4,166.6,140,167.2,138.8,167z"
		pathData71 := "M210.1,121c1.9,0.5,2.4,1.4,1.9,3.1C210.6,123.5,210.4,122.4,210.1,121z"
		pathData72 := "M163.7,146.6c0.2,1,0.6,2-0.3,2.8c-0.1,0.1-0.9-0.2-0.9-0.4C162.3,148.2,162.3,147.3,163.7,146.6z"
		pathData73 := "M123.4,155.8c0-1.4,0-3,2.1-3.3C124.7,153.8,124.1,154.8,123.4,155.8z"
		pathData74 := "M211.4,115c-3.3,0.3-1.9-1.6-2.5-3.7C210.1,113.1,210.7,114,211.4,115z"
		pathData75 := "M193.3,117.3c0.6,0,1.4,0,1.2,0.9c-0.1,0.5-0.7,0.9-1.1,1.3c-0.3-0.4-0.8-0.7-1-1.1C192.1,117.6,192.6,117.4,193.3,117.3z    "
		pathData76 := "M198.1,199.4c0.4,2.4,0.9,4-1,5.3C197.3,203.1,197.6,201.8,198.1,199.4z"
		pathData77 := "M214.8,198.4c-0.6-0.8,0.1-2-1.1-2.6c0.5-0.2,0.8-0.4,1-0.3C216.3,196.4,216.6,197.3,214.8,198.4z"
		pathData78 := "M186.3,176.5c0-1.4,0-2.9,0-4.3c0.1,0,0.3,0,0.4,0c0,1.5,0,3.1,0,4.6C186.5,176.7,186.4,176.6,186.3,176.5z"
		pathData79 := "M154.2,139.1c-0.6-2.1-0.4-3.6,2.4-3c0,0-0.1-0.1,0,0c-1.3,0.6-1.7,1.9-2.4,2.9C154.1,139,154.2,139.1,154.2,139.1z"
		pathData80 := "M142.9,106.2c-0.8,0.3-1.3,1.6-2.4,1c-0.1-0.1-0.3-0.5-0.2-0.7C140.9,105.3,141.9,106,142.9,106.2z"
		pathData81 := "M190.8,184.8c-0.1,1.1-0.2,2.2-0.4,3.3C190,186.9,189.7,185.8,190.8,184.8z"
		pathData82 := "M141.5,137.8c0.2,0.8,0.6,1.7-0.2,2.4c-0.2,0.1-0.8-0.1-0.9-0.2c-0.4-1.1,0.6-1.5,1.2-2L141.5,137.8z"
		pathData83 := "M141.6,137.9c0.8-0.5,0.5-1.9,2.4-2.3c-0.8,1.2-0.2,3-2.5,2.2C141.5,137.8,141.6,137.9,141.6,137.9z"
		pathData84 := "M97.1,163.1c0.7,1.6-0.4,2.1-1.3,2.6c-0.3-0.2-0.5-0.4-0.8-0.7c0.8-0.6,1.5-1.2,2.3-1.8C97.2,163.2,97.1,163.1,97.1,163.1    z"
		pathData85 := "M177.6,106.9c-0.7,1.5-1.7,1.2-3.2,0.9C175.9,107.4,176.7,107.2,177.6,106.9z"
		pathData86 := "M194.7,109.4c-0.7,0.1-1.1,0.2-1.5,0.2c-0.6,0-0.9-0.3-0.8-0.9c0.1-0.6,0.5-0.8,1-0.5    C193.8,108.5,194.1,108.9,194.7,109.4z"
		pathData87 := "M127.4,146.7c-0.2-0.8-0.6-1.7,0.2-2.4c0.1-0.1,0.9,0.1,0.9,0.2c0.4,1.1-0.6,1.5-1.2,2.1    C127.3,146.6,127.4,146.7,127.4,146.7z"
		pathData88 := "M129.6,117.5c-0.9,0.5-1.8,1-3.1,1.7C127.2,116.9,128.5,117.3,129.6,117.5C129.7,117.6,129.6,117.5,129.6,117.5z"
		pathData89 := "M185.9,199.2c-1.9-1.2-0.4-1.9-0.3-3C186.4,197.3,186.1,198.2,185.9,199.2z"
		pathData90 := "M167,177.9c0.7,0.1,0.9,0.6,0.6,1.1c-0.2,0.4-0.7,0.7-1,1c-0.1-0.4-0.3-0.8-0.4-1.2C166.2,178.1,166.5,177.9,167,177.9z"
		pathData91 := "M154.3,146c-0.4-1.2-0.5-2.3,0.8-3.3C154.8,143.8,154.5,144.8,154.3,146C154.2,145.9,154.3,146,154.3,146z"
		pathData92 := "M158.4,238.1c-1-5.4-1.8-10.7-0.8-16.2C157.1,227.4,160.8,232.6,158.4,238.1z"
		pathData93 := "M163.2,225.7c0-3.5,0-7.1,0-10.6c0-0.9,0.6-1.8,1.1-1.8c0.9,0.1,1.6,1,1.2,2C164.1,218.7,164.9,222.4,163.2,225.7z"
		pathData94 := "M97.8,204.9c2.2-0.8,4,0.5,5.9,1.1c0.4,1.5-1.3,0.8-1.4,1.8c-0.2,2.4-1.7,1.5-1.7,0.3c-0.1-2-2.2-1.9-2.7-3.2    C97.9,204.8,97.8,204.9,97.8,204.9z"
		pathData95 := "M113.8,237.4c5.3,2.1,5.3,2.1,7.7,0.5C120,241.7,117.9,241.7,113.8,237.4z"
		pathData96 := "M163.3,237.1c0-3.2,0-6.5,0-9.7c0.1,0,0.3,0,0.4,0c0,3.2,0,6.5,0,9.7C163.6,237.1,163.4,237.1,163.3,237.1z"
		pathData97 := "M117.2,236.8c0.2-1.9,1.2-3,2-4.4C120,234.4,119.2,235.7,117.2,236.8z"
		pathData98 := "M158.5,215.8c0,2.1,0,3.4,0,4.7C157.2,219.2,157.2,219.2,158.5,215.8z"
		pathData99 := "M138.9,217.2c-0.7-1.3,0.5-2.1,0.9-3.2c0.9,1.6-0.1,2.3-1,3.1C138.8,217.1,138.9,217.2,138.9,217.2z"
		pathData100 := "M112.9,115c-4.2-1.4-8.5-2-12.4-3.5c-4.4-1.7-8.1-4.9-9.4-10c-0.6-2.6-3.5-3-5.1-4.8c-0.3-0.4-1.2-0.4-0.8-1.8    c3.3,2.2,6.4,4.3,9.5,6.5c-0.4,1.1-1.9,0-2,1.2c-0.2,1.1,1,0.8,1.5,1.3c1.8,1.9,4.4,2,7.2,3.2c-1.6,0.2-2.7,0.3-4.5,0.4    C102.1,111.5,108.4,111.3,112.9,115z"
		pathData101 := "M100.6,129.3c0.7-0.5,1.4-1,2-1.6c1.2-1.3,2.4-1.7,4-0.7c0.4,0.3,0.8,1,0.2,1.1c-1.6,0.3-2,2.7-4.1,1.9    C102,129.7,101.3,129.5,100.6,129.3z"
		pathData102 := "M207.1,43.2c-2.3,1.1-3.8,2.8-6.4,1.9C202.7,44,204.4,43,207.1,43.2z"
		pathData103 := "M122.6,59.9c-0.2-1.3,0.3-2.2,1.7-2.3c0.1,0,0.5,0.3,0.5,0.5c-0.1,1.4-1,1.9-2.3,1.7C122.5,59.8,122.6,59.9,122.6,59.9z"
		pathData104 := "M177.4,67.7c-0.4,1.1-3,0.9-1.6,2.9c0.4,0.6,0.6,1.2-0.3,1.4c-0.7,0.2-1.2-0.4-1.4-1    C173,69.1,174.3,67.7,177.4,67.7z"
		canvas.Path(pathData0, "fill:"+lineColor)
		canvas.Path(pathData1, "fill:"+hexCode)
		canvas.Path(pathData2, "fill:"+hexCode)
		canvas.Path(pathData3, "fill:"+hexCode)
		canvas.Path(pathData4, "fill:"+hexCode)
		canvas.Path(pathData5, "fill:"+hexCode)
		canvas.Path(pathData6, "fill:"+hexCode)
		canvas.Path(pathData7, "fill:"+hexCode)
		canvas.Path(pathData8, "fill:"+hexCode)
		canvas.Path(pathData9, "fill:"+hexCode)
		canvas.Path(pathData10, "fill:"+hexCode)
		canvas.Path(pathData11, "fill:"+hexCode)
		canvas.Path(pathData12, "fill:"+hexCode)
		canvas.Path(pathData13, "fill:"+hexCode)
		canvas.Path(pathData14, "fill:"+hexCode)
		canvas.Path(pathData15, "fill:"+hexCode)
		canvas.Path(pathData16, "fill:"+hexCode)
		canvas.Path(pathData17, "fill:"+hexCode)
		canvas.Path(pathData18, "fill:"+hexCode)
		canvas.Path(pathData19, "fill:"+hexCode)
		canvas.Path(pathData20, "fill:"+hexCode)
		canvas.Path(pathData21, "fill:"+hexCode)
		canvas.Path(pathData22, "fill:"+hexCode)
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
		canvas.Path(pathData81, "fill:"+lineColor)
		canvas.Path(pathData82, "fill:"+lineColor)
		canvas.Path(pathData83, "fill:"+lineColor)
		canvas.Path(pathData84, "fill:"+lineColor)
		canvas.Path(pathData85, "fill:"+lineColor)
		canvas.Path(pathData86, "fill:"+lineColor)
		canvas.Path(pathData87, "fill:"+lineColor)
		canvas.Path(pathData88, "fill:"+lineColor)
		canvas.Path(pathData89, "fill:"+lineColor)
		canvas.Path(pathData90, "fill:"+lineColor)
		canvas.Path(pathData91, "fill:"+lineColor)
		canvas.Path(pathData92, "fill:"+lineColor)
		canvas.Path(pathData93, "fill:"+lineColor)
		canvas.Path(pathData94, "fill:"+lineColor)
		canvas.Path(pathData95, "fill:"+lineColor)
		canvas.Path(pathData96, "fill:"+lineColor)
		canvas.Path(pathData97, "fill:"+lineColor)
		canvas.Path(pathData98, "fill:"+lineColor)
		canvas.Path(pathData99, "fill:"+lineColor)
		canvas.Path(pathData100, "fill:"+lineColor)
		canvas.Path(pathData101, "fill:"+lineColor)
		canvas.Path(pathData102, "fill:"+lineColor)
		canvas.Path(pathData103, "fill:"+lineColor)
		canvas.Path(pathData104, "fill:"+hexCode)
	},
}