import React from 'react'
import { StyleProp, ViewStyle } from 'react-native'
import { LinearGradient } from 'expo-linear-gradient'


export type FrothyGradientProps = {
	children: React.ReactNode
	style?: StyleProp<ViewStyle>
}

export default class FrothyGradient extends React.Component<FrothyGradientProps, {}> {
	render() {
		const { style, children } = this.props
			return (
			<LinearGradient
			style={style}
			colors={['#FF9800', '#F44336']}
			start={{ x: 1, y: 0 }}
			end={{ x: 0.2, y: 0 }}
		>
				{children}
			</LinearGradient>
		)
	}
}
