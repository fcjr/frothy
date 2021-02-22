import React, { useState, useEffect } from 'react'
import { StyleSheet, View } from 'react-native'
import { Text, Button } from 'react-native-elements'
import { BarCodeScanner, BarCodeEvent } from 'expo-barcode-scanner'
import { BottomTabNavigationProp } from '@react-navigation/bottom-tabs'

import { MainTabsParamList } from './MainScreen'

type QRScreenNavigationProp = BottomTabNavigationProp<
	MainTabsParamList,
	'QR'
>

type QRScreenProps = {
	navigation: QRScreenNavigationProp
}

export default function QRScreen({}: QRScreenProps) {
	const [hasPermission, setHasPermission] = useState<boolean | null>(null)
	const [scanned, setScanned] = useState<boolean>(false)

	useEffect(() => {
		(async () => {
		const { status } = await BarCodeScanner.requestPermissionsAsync()
		setHasPermission(status === 'granted')
		})()
	}, [])

	const handleBarCodeScanned = ({ type, data }: BarCodeEvent) => {
		if (type !== 'org.iso.QRCode') { return }
		setScanned(true)
		alert(`Bar code with type ${type} and data ${data} has been scanned!`)
	}

	if (hasPermission === null) {
		return <Text>Requesting for camera permission</Text>
	}
	if (hasPermission === false) {
		return <Text>No access to camera</Text>
	}

	return (
		<View style={styles.container}>
		<BarCodeScanner
			onBarCodeScanned={scanned ? undefined : handleBarCodeScanned}
			style={StyleSheet.absoluteFillObject}
		/>
		{scanned && <Button title={'Tap to Scan Again'} onPress={() => setScanned(false)} />}
		</View>
	)
}

const styles = StyleSheet.create({
	container: {
		flexGrow: 1
	},
})
