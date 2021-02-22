import React, { useState } from 'react'
import { StyleSheet, Animated, View } from 'react-native'
import { ListItem } from 'react-native-elements'
import TouchableScale from 'react-native-touchable-scale'
import { CountdownCircleTimer } from 'react-native-countdown-circle-timer'
import FrothyGradient from '../components/FrothyGradient'

import { TOTP, genTOTP } from '../utils/otp'

type CodeProps = {
	uid:    string
	issuer: string
	name:   string
	secret: string
}

const updateTOTP = (secret: string, nextUpdate: Date, setTotp: React.Dispatch<React.SetStateAction<TOTP>>) => {
	setTimeout(() => {
		const totp = genTOTP(secret)
		setTotp(genTOTP(secret))
		updateTOTP(secret, totp.expiresAt, setTotp)
	}, nextUpdate.getTime() - Date.now())
}

export default function Code({ issuer, name, secret }: CodeProps) {
	const [totp, setTotp] = useState<TOTP>(genTOTP(secret))

	// TODO sync update & countdown
	updateTOTP(secret, totp.expiresAt, setTotp)
	const initialTimeRemaining = (totp.expiresAt.getTime() - Date.now()) / 1000

	return (
		<ListItem
			containerStyle={styles.container}
			Component={TouchableScale}
			friction={90}
			tension={100}
			activeScale={0.95}
			// @ts-ignore TODO figure out how to make these types work
			ViewComponent={FrothyGradient}
		>
		<View style={styles.countdownContainer}>
			<CountdownCircleTimer
				isPlaying
				size={100}
				duration={30}
				initialRemainingTime={initialTimeRemaining}
				colors="#FFF"
				onComplete={() => {
				return [true, 0]
				}}
			>
				{({ remainingTime, animatedColor }: any) => (
				<Animated.Text
					style={{ ...styles.remainingTime, color: animatedColor }}>
					{remainingTime}
				</Animated.Text>
				)}
			</CountdownCircleTimer>
			</View>
			<ListItem.Content>
				<ListItem.Title style={styles.code}>
					{totp ? totp.code : 'loading'}
				</ListItem.Title>
				<ListItem.Subtitle style={{ color: 'white' }}>
					{`${issuer}: ${name}`}
				</ListItem.Subtitle>
			</ListItem.Content>
			<ListItem.Chevron color="white" />
		</ListItem>
	)
}

const styles = StyleSheet.create({
	container: {
		height: 150,
		borderRadius: 20,
		padding: 20,
		marginTop: 15,
		marginLeft: 10,
		marginRight: 10
	},
	countdownContainer: {
		flex: 1,
		justifyContent: 'center',
		alignItems: 'center'
	},
	remainingTime: {
		fontSize: 46,
	},
	code: {
		color: 'white',
		fontWeight: 'bold',
		fontSize: 40
	}
})
