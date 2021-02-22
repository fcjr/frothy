import React from 'react'
import { StyleSheet, View } from 'react-native'
import { BottomTabNavigationProp } from '@react-navigation/bottom-tabs'
import { MainTabsParamList } from './MainScreen'
import { ScrollView } from 'react-native-gesture-handler'

import Code from '../components/Code'

type CodesScreenNavigationProp = BottomTabNavigationProp<
	MainTabsParamList,
	'Codes'
>

type CodesScreenProps = {
	navigation: CodesScreenNavigationProp
}

export default function CodesScreen({}: CodesScreenProps) {
	return (
		<View style={styles.container}>
			<ScrollView>
				<Code uid='1' issuer='Issuer A' name='User' secret='JBSWY3DPEHPK3PXP' />
				<Code uid='2' issuer='Issuer B' name='User' secret='JBSWY3DPEHPK3PXP' />
				<Code uid='3' issuer='Issuer C' name='User' secret='JBSWY3DPEHPK3PXP' />
				<Code uid='4' issuer='Issuer D' name='User' secret='JBSWY3DPEHPK3PXP' />
				<Code uid='5' issuer='Issuer E' name='User' secret='JBSWY3DPEHPK3PXP' />
				<Code uid='6' issuer='Issuer F' name='User' secret='JBSWY3DPEHPK3PXP' />
				<Code uid='7' issuer='Issuer G' name='User' secret='JBSWY3DPEHPK3PXP' />
			</ScrollView>
		</View>
	)
}

const styles = StyleSheet.create({
	container: {
		flexGrow: 1,
	}
})
