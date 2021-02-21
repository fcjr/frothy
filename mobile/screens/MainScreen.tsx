import React, { useLayoutEffect } from 'react'
import { StyleSheet } from 'react-native'
import { StackNavigationProp } from '@react-navigation/stack'
import { getFocusedRouteNameFromRoute } from '@react-navigation/native'
import { Route } from '@react-navigation/native'
import { BottomTabBar, createBottomTabNavigator } from '@react-navigation/bottom-tabs'
import Icon from 'react-native-vector-icons/MaterialIcons'

import { RootStackParamList } from '../App'
import CodesScreen from './CodesScreen'
import QRScreen from './QRScreen'
import FrothyGradient from '../components/FrothyGradient'


function getHeaderTitle(route: Partial<Route<string, object | undefined>>) {
  // If the focused route is not found, we need to assume it's the initial screen
  // This can happen during if there hasn't been any navigation inside the screen
  // In our case, it's "Feed" as that's the first screen inside the navigator
  const routeName = getFocusedRouteNameFromRoute(route) ?? 'Codes'

  switch (routeName) {
    case 'Codes':
      return 'Frothy'
    case 'QR':
      return 'Scan a 2FA QR'
  }
}

type MainScreenNavigationProp = StackNavigationProp<
    RootStackParamList,
    'Main'
>

type MainScreenProps = {
    navigation: MainScreenNavigationProp,
    route: any
}

export type MainTabsParamList = {
    Codes: {}
    QR: {}
}

const Tabs = createBottomTabNavigator<MainTabsParamList>()

export default function MainScreen({ navigation, route }: MainScreenProps) {

    useLayoutEffect(() => {
        navigation.setOptions({
            title: getHeaderTitle(route)
        })
    }, [navigation, route])

    return (
        <Tabs.Navigator
            tabBar={(props) => {
                return (
                    <FrothyGradient>
                        <BottomTabBar
                            {...props}
                            style={{ backgroundColor: 'transparent' }}
                        />
                    </FrothyGradient>
                )
            }}
            tabBarOptions={{
                activeTintColor: 'white',
                inactiveTintColor: '#D9D9D9'
            }}
        >
            <Tabs.Screen 
                name='Codes'
                component={CodesScreen}
                options={{
                    tabBarIcon: ({ color }) => (
                        <Icon name='vpn-key' size={20} color={color} />
                      )
                }}
            />
            <Tabs.Screen
                name='QR'
                component={QRScreen}
                options={{
                    tabBarIcon: ({ color }) => (
                        <Icon name='qr-code-scanner' size={20} color={color} />
                      )
                }}
            />
        </Tabs.Navigator>
    )
}

const styles = StyleSheet.create({
    container: {},
})
